package clientCommand

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"net"
	"strings"
)

type command interface {
	Execute() (string, error)
	Kind() string
}

func GetClientCommand(strArr []string) (command command, err error) {
	commandName := strArr[0]
	switch commandName {
	case shared.NEW_AGENT:
		return newAgentCommand{Key: strArr[1], Agent: strArr[2]}, nil
	case shared.PING:
		return pingCommand{Agent: strArr[1]}, nil
	case shared.ADD_REVERSE_PROXY:
		return addReverseProxyCommand{strArr: strArr}, nil
	case shared.DELETE_REVERSE_PROXY:
		return deleteReverseProxyCommand{strArr: strArr}, nil
	case shared.SAVE_AGENT_PROFILE:
		return saveAgentProfileCommand{strArr: strArr}, nil
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
