package clientCommand

import (
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/sessionInfo"
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type deleteReverseProxyCommand struct {
	strArr []string
}

// delete_reverse_proxy agentName commandName
func (cmd deleteReverseProxyCommand) Execute() (string, error) {
	agent := cmd.strArr[1]
	info := sessionInfo.GetAgentInfo(agent)
	if info == nil {
		return "", errors.New("no agent subscribed to server")
	}
	proxyName := cmd.strArr[2]

	reqMsg := fmt.Sprintf("%s %s %s", shared.DELETE_REVERSE_PROXY, agent, proxyName)
	msg, err := sendTcpMessage(info.Port, reqMsg)
	if err != nil {
		return "", err
	}
	msgSpl := strings.Split(msg, " ")
	if msgSpl[1] != "200" {
		return "", errors.New("cannot delete2 reverse proxy")
	}
	info.DeleteReverseProxy(proxyName)

	return "", nil
}

func (cmd deleteReverseProxyCommand) Kind() string {
	return shared.DELETE_REVERSE_PROXY
}
