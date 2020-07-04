package web

import "net"

type webRouter struct {
	hostToProxy map[string]net.Conn
}

var wr *webRouter

func (wr *webRouter) GetRoute(host string) net.Conn {
	return wr.hostToProxy[host]
}

func (wr *webRouter) AddRoute(host string, conn net.Conn) {
	wr.hostToProxy[host] = conn
}

func (wr *webRouter) RemoveRoute(host string) {
	delete(wr.hostToProxy, host)
}

func GetRouter() *webRouter {
	if wr == nil {
		wr = &webRouter{}
	}
	return wr
}
