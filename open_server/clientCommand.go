package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"net"
	"strings"
)

type command interface {
	Execute(server *ApiServer, s ssh.Session) error
}

type newAgentCommand struct {
	Key  string
	Agent string
}

func newCommandNewAgent(strArr []string) *newAgentCommand {
	return &newAgentCommand{
		Key:  strArr[1],
		Agent: strArr[2],
	}
}

func (n newAgentCommand) Execute(server *ApiServer, s ssh.Session) error {
	io.WriteString(s, "new_agent true nil\n")
	server.agentSession[n.Key] = AgentInfo{
		Agent:    n.Agent,
	}
	return nil
}

type pingCommand struct {
	Agent string
}

func newPingCommand(strArr []string) *pingCommand {
	return &pingCommand{
		Agent: strArr[1],
	}
}

func (n pingCommand) Execute(server *ApiServer, s ssh.Session) error {
	success := false
	for _,v := range server.agentSession{
		agentSession := v.Session
		if v.Session == nil || agentSession.User() != n.Agent{
			continue
		}
		localConn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", v.Port))
		if err != nil {
			return err
		}
		_, err = localConn.Write([]byte("ping\n"))
		if err != nil {
			return err
		}
		data, err := bufio.NewReader(localConn).ReadString('\n')
		if err != nil {
			return err
		}
		msg := strings.TrimSpace(data)
		if msg != "pong"{
			return errors.New("response different then pong")
		}
		success = true
	}
	if !success {
		return errors.New("no success on response")
	}
	io.WriteString(s, "ping true nil\n")
	return nil
}

func getClientCommand(strArr []string) (command command, err error) {
	commandName := strArr[0]
	switch commandName {
	case "new_agent":
		return newCommandNewAgent(strArr), nil
	case "ping":
		return newPingCommand(strArr), nil
	}
	return nil, errors.New("cannot find command")
}
