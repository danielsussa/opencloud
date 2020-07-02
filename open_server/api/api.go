package main

import (
	"github.com/gliderlabs/ssh"
	"log"
)

type Config struct {
	ServerClientPort     string
	ServerClientPassword string
	ServerAgentPort      string
}

type ApiServer struct {
	Config Config

	// ssh keys
	agentSession map[ssh.PublicKey]AgentInfo
}

type AgentInfo struct {
	Name    string
	Session *ssh.Session
}

func (apiServer *ApiServer) Start(config Config) {
	apiServer.agentSession = make(map[ssh.PublicKey]AgentInfo)
	errChan := make(chan error, 2)
	apiServer.Config = config
	apiServer.serverAgentHandler(config, errChan)
	apiServer.serverClientHandler(config, errChan)
	select {
	case err := <-errChan:
		if err != nil {
			log.Fatal(err)
		}
	}
}

var password = "custom_password"

func main() {

	config := Config{
		ServerClientPassword: password,
		ServerClientPort:     ":2222",
		ServerAgentPort:      ":9999",
	}
	server := ApiServer{}

	server.Start(config)
}
