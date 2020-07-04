package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type addReverseProxy struct {
	flags flags
}

// add_reverse_proxy agentName 8080
func (rp addReverseProxy) Request() string {
	return fmt.Sprintf("%s %s %s", shared.ADD_REVERSE_PROXY, *rp.flags.Agent, *rp.flags.Port)
}

func (rp addReverseProxy) Response(strArr []string) string {
	if strArr[1] != "200"{
		return fmt.Sprintf("error creating proxy: %s", strings.Join(strArr[2:]," "))
	}
	return fmt.Sprintf("sucessfull create new proxy")
}
