package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	info "github.com/danielsussa/opencloud/open_server/info"
	"github.com/danielsussa/opencloud/shared"
	"github.com/gliderlabs/ssh"
	"github.com/phayes/freeport"
	"io"
)

type connectAgentCommand struct {
	Agent string
}

func (n connectAgentCommand) Kind() string {
	return shared.CONNECT_AGENT
}

func newConnectAgentCommand() *connectAgentCommand {
	return &connectAgentCommand{}
}

func (n connectAgentCommand) Execute(apiServer *ApiServer, s ssh.Session) (string, error) {
	keyEncoded := base64.StdEncoding.EncodeToString(s.PublicKey().Marshal())
	port, err := freeport.GetFreePort()
	if err != nil {
		return "", err
	}

	info.AddAgentInfo(s.User(), keyEncoded, port, s)
	_, err = io.WriteString(s, fmt.Sprintf("connect_agent %d", port))
	if err != nil {
		return "", err
	}
	return "", nil
}

func getAgentCommand(strArr []string) (command command, err error) {
	commandName := strArr[0]
	switch commandName {
	case shared.CONNECT_AGENT:
		return newConnectAgentCommand(), nil
	}
	return nil, errors.New("cannot find command")
}
