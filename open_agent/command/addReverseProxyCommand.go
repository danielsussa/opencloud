package command

import (
	"fmt"
	sshUtils "github.com/danielsussa/opencloud/open_agent/ssh"
	"github.com/danielsussa/opencloud/shared"
	"io"
	"log"
	"net"
	"strconv"
)

type addReverseProxyCommand struct {
	cmd []string
}


func (p addReverseProxyCommand) Kind() string {
	return shared.ADD_REVERSE_PROXY
}

// add_reverse_proxy proxyName 8080 52738
func (p addReverseProxyCommand) Execute() (string, error) {
	// read message
	user := p.cmd[1]
	localPort , err := strconv.Atoi(p.cmd[2])
	if err != nil {
		return "",err
	}
	remotePort , err := strconv.Atoi(p.cmd[3])
	if err != nil {
		return "",err
	}
	keyPair := sshUtils.GetOrGenerateRsaKeyGen()
	remoteHost := fmt.Sprintf("localhost:%d", remotePort)
	localHost := fmt.Sprintf("localhost:%d", localPort)

	err = reverseProxy(remoteHost, localHost, user, keyPair.Private)
	if err != nil {
		return "",err
	}
	return "",nil
}

func reverseProxy(remoteHost, localHost,user string, key []byte) error{
	remoteCli, err := sshUtils.SshClient(remoteHost, user,key)
	if err != nil {
		return err
	}
	// Listen on remote server port
	listener, err := remoteCli.Listen("tcp", remoteHost)
	if err != nil {
		return err
	}

	go func(){
		// handle incoming connections on reverse forwarded tunnel
		for {
			client, err := listener.Accept()
			if err != nil {
				log.Fatalln(err)
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
