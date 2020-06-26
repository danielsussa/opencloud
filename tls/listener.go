package tlsListener

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

type listener struct {
	serverTls *tls.Config
	k func(conn net.Conn)
}

func (l listener)Listen(port string){
	listener, err := tls.Listen("tcp", port, l.serverTls)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("listen to port %s", port))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go l.k(conn)
	}
}

func New(k func(conn net.Conn)) *listener{
	serverConf,_, err := certsetup()
	if err != nil {
		log.Fatal(err)
	}
	return &listener{
		serverTls: serverConf,
		k:         k,
	}
}
