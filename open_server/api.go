package main

import (
	"github.com/danielsussa/opencloud/open_server/api"
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/open_server/flags"
)

var password = "custom_password"

func main() {
	data.InitFromFile()
	config := api.Config{
		ServerClientPassword: password,
		ServerClientPort:     flags.GetClientPort(),
		ServerAgentPort:      flags.GetQgentPort(),
	}
	server := api.ApiServer{}

	server.Start(config)
}
