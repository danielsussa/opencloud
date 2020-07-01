package webProxy

import (
	"errors"
	"fmt"
	tlsListener "github.com/danielsussa/opencloud/open_server/tls"
	"github.com/danielsussa/opencloud/open_server/web"
	webRouter "github.com/danielsussa/opencloud/open_server/web_router"
	"io"
	"log"
	"net"
	"strings"
)

var currentHost string

func extractHost(chain string) (string, error) {
	log.Println(tlsListener.GetCACertificate(chain).DNSNames)
	host := tlsListener.GetCACertificate(chain).DNSNames[0]
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

	httpInfo := web.ExtractHttpInfo(conn)
	// redirect to
	if httpInfo.Host() == currentHost {
		go handleWebCommand(conn, httpInfo)
		return
	}

	router := webRouter.GetRouter()
	proxyConn := router.GetRoute(httpInfo.Host())
	if proxyConn == nil {
		return
	}

	go handleRedirectClient(proxyConn, conn)
}

func handleWebCommand(conn net.Conn, info *web.HttpInfo) {
	fmt.Println(info)
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
