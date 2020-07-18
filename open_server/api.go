package main

import (
	"github.com/danielsussa/opencloud/open_server/api"
)

var password = "custom_password"

func main() {
	config := api.Config{
		ServerClientPassword: password,
		ServerClientPort:     ":9999",
		ServerAgentPort:      ":2222",
	}
	server := api.ApiServer{}

	server.Start(config)
}
