package clientCommand

import (
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/shared"
)

type saveAgentProfileCommand struct {
	strArr []string
}

func (n saveAgentProfileCommand) Kind() string {
	return shared.SAVE_AGENT_PROFILE
}

func (n saveAgentProfileCommand) Execute() (string, error) {
	agent := n.strArr[1]
	data.SaveAgentProfile(agent)
	return "", nil
}
