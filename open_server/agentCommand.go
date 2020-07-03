package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/phayes/freeport"
	"io"
)


type connectAgentCommand struct {
	Agent string
}

func newConnectAgentCommand() *connectAgentCommand {
	return &connectAgentCommand{
	}
}

func (n connectAgentCommand) Execute(apiServer *ApiServer, s ssh.Session) error {
	keyEncoded := base64.StdEncoding.EncodeToString(s.PublicKey().Marshal())
	port, err := freeport.GetFreePort()
	if err != nil {
		return err
	}

	info := apiServer.agentSession[keyEncoded]
	info.Session = s
	info.Port = port
	apiServer.agentSession[keyEncoded] = info
	_, err = io.WriteString(s, fmt.Sprintf("connect_agent %d", port))
	if err != nil {
		return err
	}
	return nil
}

func getAgentCommand(strArr []string) (command command, err error) {
	commandName := strArr[0]
	switch commandName {
	case "connect_agent":
		return newConnectAgentCommand(), nil
	}
	return nil, errors.New("cannot find command")
}
