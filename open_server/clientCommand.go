package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/info"
	"github.com/danielsussa/opencloud/shared"
	"github.com/gliderlabs/ssh"
	"github.com/phayes/freeport"
	"net"
	"strconv"
	"strings"
)

type command interface {
	Execute(server *ApiServer, s ssh.Session) (string, error)
	Kind() string
}

type newAgentCommand struct {
	Key   string
	Agent string
}

func (n newAgentCommand) Kind() string {
	return shared.NEW_AGENT
}

func newCommandNewAgent(strArr []string) *newAgentCommand {
	return &newAgentCommand{
		Key:   strArr[1],
		Agent: strArr[2],
	}
}

func (n newAgentCommand) Execute(server *ApiServer, s ssh.Session) (string, error) {
	info.AddAgentInfoKey(n.Agent, n.Key)
	return "ok", nil
}

type pingCommand struct {
	Agent string
}

func (n pingCommand) Kind() string {
	return shared.PING
}

func newPingCommand(strArr []string) *pingCommand {
	return &pingCommand{
		Agent: strArr[1],
	}
}

func (n pingCommand) Execute(server *ApiServer, s ssh.Session) (string, error) {
	info := info.GetAgentInfo(n.Agent)
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

type addReverseProxyCommand struct {
	strArr []string
}

// add_reverse_proxy agentName 8080
func (cmd addReverseProxyCommand) Execute(server *ApiServer, s ssh.Session) (string, error) {
	agent := cmd.strArr[1]
	info := info.GetAgentInfo(agent)
	if info == nil {
		return "", errors.New("no agent subscribed to server")
	}
	localPort, err := strconv.Atoi(cmd.strArr[2])
	if err != nil {
		return "", err
	}
	remotePort, err := freeport.GetFreePort()
	if err != nil {
		return "", err
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

	fmt.Println(agent, localPort, remotePort, msg)
	return "success to add reverse proxy!", nil
}

func (a addReverseProxyCommand) Kind() string {
	return shared.ADD_REVERSE_PROXY
}

func getClientCommand(strArr []string) (command command, err error) {
	commandName := strArr[0]
	switch commandName {
	case shared.NEW_AGENT:
		return newCommandNewAgent(strArr), nil
	case shared.PING:
		return newPingCommand(strArr), nil
	case shared.ADD_REVERSE_PROXY:
		return addReverseProxyCommand{strArr: strArr}, nil
	}
	return nil, errors.New("cannot find command")
}

func sendTcpMessage(port int, msg string) (string, error) {
	localConn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return "", err
	}
	_, err = localConn.Write([]byte(fmt.Sprintf("%s\n", msg)))
	if err != nil {
		return "", err
	}
	data, err := bufio.NewReader(localConn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(data), nil
}
