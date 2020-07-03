package main

import (
	"encoding/json"
	"fmt"
	"github.com/gliderlabs/ssh"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

type Config struct {
	ServerClientPort     string
	ServerClientPassword string
	ServerAgentPort      string

	SavedCommandList [][]string
}

type ApiServer struct {
	Config      Config
	CommandList [][]string

	agentSession map[string]AgentInfo // public rsa key encoded base64
}

type AgentInfo struct {
	Agent    string
	Session   ssh.Session
	Port int
}

func (apiServer *ApiServer) Start(config Config) {
	apiServer.agentSession = make(map[string]AgentInfo)
	errChan := make(chan error, 2)
	apiServer.Config = config
	apiServer.serverAgentHandler(config, errChan)
	apiServer.serverClientHandler(config, errChan)
	apiServer.gracefullTerminate()
	select {
	case err := <-errChan:
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (apiServer *ApiServer) gracefullTerminate() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case sig := <-c:
			fmt.Printf("Got %s signal. Aborting...\n", sig)
			if apiServer.CommandList == nil {
				return
			}
			b, err := json.Marshal(apiServer.CommandList)
			fmt.Println(err)
			ioutil.WriteFile("data/commands.json", b, 0644)
			os.Exit(1)
		}
	}()
}

var password = "custom_password"

func main() {
	config := Config{
		ServerClientPassword: password,
		ServerClientPort:     ":9999",
		ServerAgentPort:      ":2222",
	}
	server := ApiServer{}

	server.Start(config)
}
