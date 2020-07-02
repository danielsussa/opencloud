package main

import (
	"encoding/base64"
	"errors"
	"github.com/gliderlabs/ssh"
)

type command interface {
	Kind() string
	Execute(server *ApiServer, s ssh.Session) error
}

type newAgent struct {
	Key  string
	Name string
}

func newCommandNewAgent(strArr []string) *newAgent {
	return &newAgent{
		Key:  strArr[1],
		Name: strArr[2],
	}
}

func (n newAgent) Execute(server *ApiServer, s ssh.Session) error {
	b, err := base64.StdEncoding.DecodeString(n.Key)
	if err != nil {
		return err
	}
	parsedKey, _, _, _, err := ssh.ParseAuthorizedKey(b)
	if err != nil {
		return err
	}

	server.agentSession[parsedKey] = AgentInfo{
		Name:    n.Name,
		Session: &s,
	}
	return nil
}

func (n newAgent) Kind() string {
	return "new_agent"
}

func getCommand(strArr []string) (command command, err error) {
	commandName := strArr[0]
	switch commandName {
	case "new_agent":
		return newCommandNewAgent(strArr), nil
	}
	return nil, errors.New("cannot find command")
}
