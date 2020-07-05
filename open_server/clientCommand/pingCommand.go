package clientCommand

import (
	"errors"
	"github.com/danielsussa/opencloud/open_server/sessionInfo"
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type pingCommand struct {
	Agent string
}

func (n pingCommand) Kind() string {
	return shared.PING
}

func (n pingCommand) Execute() (string, error) {
	info := sessionInfo.GetAgentInfo(n.Agent)
	if info == nil {
		return "", errors.New("no agent subscribed to server")
	}

	msg, err := sendTcpMessage(info.Port, "ping")
	if err != nil {
		return "", err
	}
	pongMsg := strings.Split(msg, " ")[2]
	if pongMsg != "pong" {
		return "", errors.New("response different then pong")
	}
	return "ping 200 nil", nil
}
