package command

import (
	"github.com/danielsussa/opencloud/shared"
	"net"
	"strings"
)

type ApiCommand interface {
	Execute(conn net.Conn)error
}

func CommandHandler(cmd string) ApiCommand {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[0] {
	case shared.PING:
		return pingCommand{}
	case shared.ADD_REVERSE_PROXY:
		return addReverseProxyCommand{cmd: cmdArr}
	}
	return nil
}
