package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type saveAgentProfileCommand struct {
	flags flags
}

// ping agentName
func (ping saveAgentProfileCommand) Request() string {
	return fmt.Sprintf("%s %s", shared.SAVE_AGENT_PROFILE, *ping.flags.Agent)
}

func (ping saveAgentProfileCommand) Response(strArr []string) string {
	if strArr[1] != "200"{
		return fmt.Sprintf("error to save agent: %s", strings.Join(strArr[2:]," "))
	}
	return fmt.Sprintf("sucessfull save agent")
}
