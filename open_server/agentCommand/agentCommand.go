package agentCommand

import (
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/shared"
	"github.com/gliderlabs/ssh"
	"io"
)

type agentCommand interface {
	Execute(s ssh.Session) (string, error)
	Kind() string
}

type connectAgentCommand struct {
	Agent string
}

func (n connectAgentCommand) Kind() string {
	return shared.CONNECT_AGENT
}

func (n connectAgentCommand) Execute(s ssh.Session) (string, error) {
	agentData := data.GetAgentData(s.User())

	port, err := agentData.GetPort()
	if err != nil {
		return "", err
	}

	commands := fmt.Sprintf("%s %d\n", shared.CONNECT_AGENT, port)

	if agentData != nil {
		for rpName, rp := range agentData.ReverseProxy {
			// add_reverse_proxy myagent-1 proxyName 1323 52738
			commands += fmt.Sprintf("%s %s %s %d %d\n", shared.ADD_REVERSE_PROXY, s.User(), rpName, rp.LocalPort, rp.RemotePort)
		}
	}

	_, err = io.WriteString(s, commands)
	if err != nil {
		return "", err
	}
	return "", nil
}

func GetAgentCommand(strArr []string) (agentCommand, error) {
	commandName := strArr[0]
	switch commandName {
	case shared.CONNECT_AGENT:
		return &connectAgentCommand{}, nil
	}
	return nil, errors.New("cannot find command")
}
