package clientCommand

import (
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/sessionInfo"
	"github.com/danielsussa/opencloud/shared"
	"github.com/phayes/freeport"
	"strconv"
	"strings"
)

type addReverseProxyCommand struct {
	strArr []string
}

// add_reverse_proxy agentName commandName 8080
func (cmd addReverseProxyCommand) Execute() (string, error) {
	agent := cmd.strArr[1]
	info := sessionInfo.GetAgentInfo(agent)
	if info == nil {
		return "", errors.New("no agent subscribed to server")
	}
	proxyName := cmd.strArr[2]
	localPort, err := strconv.Atoi(cmd.strArr[3])
	if err != nil {
		return "", err
	}
	remotePort, err := freeport.GetFreePort()
	if err != nil {
		return "", err
	}
	if info.HasReverseProxy(localPort) {
		return "", errors.New("reverse proxy already exist")
	}

	reqMsg := fmt.Sprintf("%s %s %d %d", shared.ADD_REVERSE_PROXY, agent, localPort, remotePort)
	msg, err := sendTcpMessage(info.Port, reqMsg)
	if err != nil {
		return "", err
	}
	msgSpl := strings.Split(msg, " ")
	if msgSpl[1] != "200" {
		return "", errors.New("cannot operate reverse proxy")
	}
	info.AddReverseProxy(proxyName, localPort, remotePort)
	return "success to add reverse proxy!", nil
}

func (a addReverseProxyCommand) Kind() string {
	return shared.ADD_REVERSE_PROXY
}
