package api

import (
	"fmt"
	errorhandler "github.com/danielsussa/opencloud/open_server/error"
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
	apiServer.Config = config
	apiServer.serverAgentHandler(config)
	apiServer.serverClientHandler(config)
	apiServer.tlsHandler()
	//apiServer.httpHandler()

	apiServer.gracefullTerminate()

	errorhandler.ListenForErrors()

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
