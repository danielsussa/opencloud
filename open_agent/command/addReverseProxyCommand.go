package command

import (
	"net"
)

type addReverseProxyCommand struct {
	cmd []string
}

func (p addReverseProxyCommand)Execute(conn net.Conn)error{
	// exemple: add_reverse_proxy proxyName 8080 52738
	_, err := conn.Write([]byte("pong\n"))
	if err != nil {
		return err
	}
	return nil
}
