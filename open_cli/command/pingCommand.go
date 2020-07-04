package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
)

type pingCommand struct {
	flags flags
}

// ping agentName
func (ping pingCommand) Request() string {
	return fmt.Sprintf("%s %s", shared.PING, *ping.flags.Agent)
}

func (ping pingCommand) Response(strArr []string) string {
	return fmt.Sprintf("sucessfull ping to host")
}
