package clientCommand

import (
	"github.com/danielsussa/opencloud/open_server/sessionInfo"
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
	sessionInfo.AddAgentInfoKey(n.Agent, n.Key)
	return "ok", nil
}
