package tcpListener

import (
	"crypto/tls"
	"github.com/danielsussa/opencloud/open_server/web"
	"log"
	"net"
)

func New(key web.RsaKeyPair, port string, k func(conn net.Conn)) {
	cer, err := tls.X509KeyPair(key.Key, key.Public)
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	l, err := tls.Listen("tcp", port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("listen tcp on port ", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go k(conn)
	}
}
