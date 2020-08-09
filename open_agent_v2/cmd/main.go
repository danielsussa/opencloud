package main

import (
	applicatioLayer "github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer"
	"github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer/push"
	domainLayer"github.com/danielsussa/opencloud/open_agent_v2/internal/domain_layer"
	"github.com/danielsussa/opencloud/open_agent_v2/internal/domain_layer/domain"
	"github.com/danielsussa/opencloud/open_agent_v2/internal/domain_layer/ipush"
	"log"
)

func main(){
	err := applicatioLayer.Setup()
	if err != nil {
		log.Fatal(err)
	}
	domainLayer.Setup(domainLayer.Config{
		Push: ipush.Push{
			LoadPresets: push.LoadPresets,
		},
	})

	domain.LoadPresetsDomain()
}
