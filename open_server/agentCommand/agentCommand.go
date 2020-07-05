package agentCommand

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/open_server/sessionInfo"
	"github.com/danielsussa/opencloud/shared"
	"github.com/gliderlabs/ssh"
	"github.com/phayes/freeport"
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
	keyEncoded := base64.StdEncoding.EncodeToString(s.PublicKey().Marshal())
	port, err := freeport.GetFreePort()
	if err != nil {
		return "", err
	}

	sessionInfo.AddAgentInfo(s.User(), keyEncoded, port, s)

	commands := fmt.Sprintf("%s %d\n", shared.CONNECT_AGENT, port)
	agentData := data.GetAgentData(s.User())
	if agentData != nil {
		for rpName, rp := range agentData.ReverseProxy {
			rpFreePort, err := freeport.GetFreePort()
			if err != nil {
				return "", err
			}
			// add_reverse_proxy myagent-1 proxyName 1323 52738
			commands += fmt.Sprintf("%s %s %s %d %d\n", shared.ADD_REVERSE_PROXY, s.User(), rpName, rp.Port, rpFreePort)
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
