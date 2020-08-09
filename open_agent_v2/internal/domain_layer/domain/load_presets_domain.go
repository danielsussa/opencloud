package domain

import "github.com/danielsussa/opencloud/open_agent_v2/internal/domain_layer/ipush"

func LoadPresetsDomain(){
	ipush.Get().LoadPresets()
}
