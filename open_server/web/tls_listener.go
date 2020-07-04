package web

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

type listener struct {
	serverTls *tls.Config
	k         func(conn net.Conn)
}

func (l listener) Listen(port string) {
	listener, err := tls.Listen("tcp", port, l.serverTls)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("listen to port %s", port))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go l.k(conn)
	}
}

func New(chain string, pKey string, k func(conn net.Conn)) *listener {
	var err error
	conf := tls.Config{}
	conf.Certificates = make([]tls.Certificate, 1)
	conf.Certificates[0], err = tls.LoadX509KeyPair(chain, pKey)
	if err != nil {
		log.Fatal(err)
	}
	return &listener{
		serverTls: &conf,
		k:         k,
	}
}
