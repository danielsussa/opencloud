package web

import (
	"bufio"
	"errors"
	"fmt"
	tlsListener "github.com/danielsussa/opencloud/open_server/tls"
	webRouter "github.com/danielsussa/opencloud/open_server/web_router"
	"io"
	"log"
	"net"
	"strings"
)

var currentHost string

func extractHost(chain string) (string, error) {
	host := tlsListener.GetCertificate(chain).DNSNames[0]
	if !strings.HasPrefix(host, "*.") {
		return "", errors.New("cannot extract wildcard from cert")
	}
	return host[2:], nil
}

func ServeHttps(chain, key, port string) {
	var err error
	currentHost, err = extractHost(chain)
	if err != nil {
		log.Fatal(err)
	}
	tlsListener.New(chain, key, listenerConn).
		Listen(port)
}

func listenerConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	host := ""
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		if strings.Contains(msg, "Host: ") {
			host = msg[6 : len(msg)-len(currentHost)-3]
			break
		}
	}
	if host == "" {
		return
	}

	router := webRouter.GetRouter()
	proxyConn := router.GetRoute(host)
	if proxyConn == nil {
		return
	}

	go handleClient(proxyConn, conn)
}

func handleClient(remoteConn net.Conn, localConn net.Conn) {
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
