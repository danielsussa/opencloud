package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
)

type newAgentCommand struct {
	flags flags
}

// new_agent d92d2f2f2f28f28f28f agentName
func (newAgent newAgentCommand) Request() string {
	return fmt.Sprintf("%s %s %s", shared.NEW_AGENT, *newAgent.flags.Key, *newAgent.flags.Agent)
}

func (newAgent newAgentCommand) Response(strArr []string) string {
	return fmt.Sprintf("sucessfull creating new agent")
}
