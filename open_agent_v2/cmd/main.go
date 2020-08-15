package main

import (
	applicatioLayer "github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer"
	"log"
)

func main(){
	err := applicatioLayer.Start()
	if err != nil {
		log.Fatal(err)
	}

}
