package command

import (
	"github.com/danielsussa/opencloud/open_server/tcp"
	tlsListener "github.com/danielsussa/opencloud/open_server/tls"
	"net"
)

var hasInstance = false

func Run(keyPair tlsListener.RsaKeyPair, port string) {
	if hasInstance {
		return
	}
	hasInstance = true
	go tcpListener.New(keyPair, port, listenerConn)
}

func listenerConn(conn net.Conn) {
}
