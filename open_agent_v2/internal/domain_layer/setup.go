package domain_layer

import "github.com/danielsussa/opencloud/open_agent_v2/internal/domain_layer/ipush"

type Config struct {
	Push ipush.Push
}

func Setup(c Config){
	ipush.Setup(c.Push)
}
