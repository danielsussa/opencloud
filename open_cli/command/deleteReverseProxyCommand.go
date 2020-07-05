package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type deleteReverseProxy struct {
	flags flags
}

// add_reverse_proxy agentName proxyName 8080
func (rp deleteReverseProxy) Request() string {
	return fmt.Sprintf("%s %s %s", shared.DELETE_REVERSE_PROXY, *rp.flags.Agent,*rp.flags.Name)
}

func (rp deleteReverseProxy) Response(strArr []string) string {
	if strArr[1] != "200"{
		return fmt.Sprintf("error deleting proxy: %s", strings.Join(strArr[2:]," "))
	}
	return fmt.Sprintf("sucessfull delete proxy")
}
