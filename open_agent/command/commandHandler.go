package command

import (
	"github.com/danielsussa/opencloud/shared"
	"strings"
)

type ApiCommand interface {
	Execute()(string,error)
	Kind()string
}

func CommandHandler(cmd string) ApiCommand {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[0] {
	case shared.PING:
		return pingCommand{}
	case shared.ADD_REVERSE_PROXY:
		return addReverseProxyCommand{cmd: cmdArr}
	case shared.DELETE_REVERSE_PROXY:
		return deleteReverseProxyCommand{cmd: cmdArr}
	}
	return nil
}
