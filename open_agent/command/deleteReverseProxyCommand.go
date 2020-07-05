package command

import (
	"github.com/danielsussa/opencloud/open_agent/reverseProxy"
	"github.com/danielsussa/opencloud/shared"
)

type deleteReverseProxyCommand struct {
	cmd []string
}


func (p deleteReverseProxyCommand) Kind() string {
	return shared.ADD_REVERSE_PROXY
}

// add_reverse_proxy agent proxyName
func (p deleteReverseProxyCommand) Execute() (string, error) {
	// read message
	proxyName := p.cmd[2]

	err := reverseProxy.DeleteReverseProxy(proxyName)
	if err != nil {
		return "",err
	}
	return "",nil
}
