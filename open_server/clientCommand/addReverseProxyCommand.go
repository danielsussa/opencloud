package clientCommand

import (
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/shared"
	"strconv"
	"strings"
)

type addReverseProxyCommand struct {
	strArr []string
}

// add_reverse_proxy agentName commandName 8080
func (cmd addReverseProxyCommand) Execute() (string, error) {
	agent := cmd.strArr[1]
	agentData := data.GetAgentData(agent)
	if agentData == nil {
		return "", errors.New("no agent subscribed to server")
	}
	proxyName := cmd.strArr[2]
	localPort, err := strconv.Atoi(cmd.strArr[3])
	if err != nil {
		return "", err
	}

	if agentData.HasReverseProxy(localPort) {
		return "", errors.New("reverse proxy already exist")
	}
	remotePort, err := data.GetNewFreeNotAllocatedPort(data.GetData())
	if err != nil {
		return "", err
	}
	reqMsg := fmt.Sprintf("%s %s %s %d %d", shared.ADD_REVERSE_PROXY, agent, proxyName, localPort, remotePort)
	msg, err := sendTcpMessage(agentData.Port, reqMsg)
	if err != nil {
		return "", err
	}
	msgSpl := strings.Split(msg, " ")
	if msgSpl[1] != "200" {
		return "", errors.New("cannot operate reverse proxy")
	}
	agentData.AddReverseProxy(proxyName, localPort, remotePort)
	return "success to add reverse proxy!", nil
}

func (a addReverseProxyCommand) Kind() string {
	return shared.ADD_REVERSE_PROXY
}
