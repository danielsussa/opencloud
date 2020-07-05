package clientCommand

import (
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/shared"
)

type newAgentCommand struct {
	Key   string
	Agent string
}

func (n newAgentCommand) Kind() string {
	return shared.NEW_AGENT
}

func (n newAgentCommand) Execute() (string, error) {
	data.GetData().NewAgent(n.Agent, n.Key)
	return "ok", nil
}
