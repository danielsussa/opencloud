package api

import (
	"crypto/tls"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/danielsussa/opencloud/open_server/flags"
	"github.com/danielsussa/opencloud/open_server/web"
	"io"
	"log"
	"net"
	"time"
)

func (apiServer *ApiServer) tlsHandler() {

	chain := flags.GetTlsPublic()
	pKey := flags.GetTlsPrivate()

	if chain == "" && pKey == "" {
		return
	}

	var err error
	conf := &tls.Config{}
	conf.Certificates = make([]tls.Certificate, 1)
	conf.Certificates[0], err = tls.LoadX509KeyPair(chain, pKey)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := tls.Listen("tcp", ":443", conf)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	httpInfo := web.ExtractHttpInfo(conn)
	port, err := data.GetData().Web.GetPort(httpInfo.Host())
	if err != nil {
		log.Println(fmt.Printf("Cannot find port: %s", err))
		conn.Close()
		return
	}
	localConn, err := net.DialTimeout("tcp", fmt.Sprintf(":%d", port), 5*time.Second)
	if err != nil {
		log.Println(fmt.Printf("Dial INTO local service error: %s", err))
		conn.Close()
		return
	}
	handleRedirectClient(localConn, conn)
}

func handleRedirectClient(remoteConn net.Conn, localConn net.Conn) {
	defer remoteConn.Close()
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(remoteConn, localConn)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy remote->local: %s", err))
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(localConn, remoteConn)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy local->remote: %s", err))
		}
		chDone <- true
	}()

	<-chDone
}
