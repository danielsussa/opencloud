package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type pingCommand struct {
	flags flags
}

// ping agentName
func (ping pingCommand) Request() string {
	return fmt.Sprintf("%s %s", shared.PING, *ping.flags.Agent)
}

func (ping pingCommand) Response(strArr []string) string {
	if strArr[1] != "200"{
		return fmt.Sprintf("error ping: %s", strings.Join(strArr[2:]," "))
	}
	return fmt.Sprintf("sucessfull ping to host")
}
