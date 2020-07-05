package reverseProxy

import (
	"errors"
	"fmt"
	sshUtils "github.com/danielsussa/opencloud/open_agent/ssh"
	"io"
	"log"
	"net"
)

var proxyMap map[string]net.Listener

func init() {
	proxyMap = make(map[string]net.Listener)
}

func DeleteReverseProxy(proxyName string)error{
	if listener, ok := proxyMap[proxyName]; ok {
		listener.Close()
		return nil
	}
	return errors.New("proxy doesnt exist")
}

func NewReverseProxy(remoteHost, localHost, user, proxyName string, key []byte) error {
	remoteCli, err := sshUtils.SshClient(remoteHost, user,key)
	if err != nil {
		return err
	}
	// Listen on remote server port
	listener, err := remoteCli.Listen("tcp", remoteHost)
	if err != nil {
		return err
	}

	proxyMap[proxyName] = listener

	go func(){
		for {
			client, err := listener.Accept()
			if err != nil {
				return
			}
			// Open a (local) connection to localEndpoint whose content will be forwarded so serverEndpoint
			local, err := net.Dial("tcp", localHost)
			if err != nil {
				log.Println(err)
				continue
			}
			handleClient(client, local)
		}
	}()
	return nil
}

func handleClient(client net.Conn, remote net.Conn) {
	defer client.Close()
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy remote->local: %s", err))
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy local->remote: %s", err))
		}
		chDone <- true
	}()

	<-chDone
}

