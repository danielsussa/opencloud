package main

import (
	"github.com/gliderlabs/ssh"
	"log"
)

type Config struct {
	ServerClientPort string
	ServerAgentPort  string
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
	errChan := make(chan error, 2)
	apiServer.serverAgentHandler(config, errChan)
	apiServer.serverClientHandler(config, errChan)
	select {
	case err := <-errChan:
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {

	config := Config{
		ServerClientPort: ":2222",
		ServerAgentPort:  ":9999",
	}
	server := ApiServer{}

	server.Start(config)
}
