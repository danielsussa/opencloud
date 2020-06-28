package command

import (
	tlsListener "github.com/danielsussa/opencloud/open_server/tls"
	"net"
)

var hasInstance = false

func Run(chain, key, port string) {
	if hasInstance {
		return
	}
	hasInstance = true
	go tlsListener.New(chain, key, listenerConn).
		Listen(port)
}

func listenerConn(conn net.Conn) {
}
