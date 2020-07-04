package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

type Config struct {
	ServerClientPort     string
	ServerClientPassword string
	ServerAgentPort      string
}

type ApiServer struct {
	Config Config
}

func (apiServer *ApiServer) Start(config Config) {
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
			os.Exit(1)
			//if apiServer.CommandList == nil {
			//	os.Exit(1)
			//}
			//b, err := json.Marshal(apiServer.CommandList)
			//fmt.Println(err)
			//ioutil.WriteFile("data/commands.json", b, 0644)
			//os.Exit(1)
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
