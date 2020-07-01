package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
)

type ReverseProxyInfo struct {
	id         string
	localPort  int
	Host       string
	RemotePort int
}

func (rp ReverseProxyInfo) Kind() CommandType {
	return ADD_REVERSE_PROXY
}

func (rp ReverseProxyInfo) RemoteString() string {
	return fmt.Sprintf("%s:%d", rp.Host, rp.RemotePort)
}

func (rp ReverseProxyInfo) LocalString() string {
	return fmt.Sprintf("%s:%d", rp.Host, rp.localPort)
}

type AddSshReverseProxy func(c Client, info ReverseProxyInfo)

func addReverseProxy(c Client, info ReverseProxyInfo) {

	key, err := ssh.ParsePrivateKey(c.rsaKeyPair.Private)
	if err != nil {
		log.Fatalln(err)
	}

	// refer to https://godoc.org/golang.org/x/crypto/ssh for other authentication types
	sshConfig := &ssh.ClientConfig{
		// SSH connection username
		User:            c.Config.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(key)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH remote server using serverEndpoint
	serverConn, err := ssh.Dial("tcp", c.Config.ParseSshHostPort(), sshConfig)
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}

	// Listen on remote server port
	listener, err := serverConn.Listen("tcp", info.RemoteString())
	if err != nil {
		log.Fatalln(fmt.Printf("Listen open port ON remote server error: %s", err))
	}
	defer listener.Close()

	// handle incoming connections on reverse forwarded tunnel
	for {
		// Open a (local) connection to localEndpoint whose content will be forwarded so serverEndpoint
		local, err := net.Dial("tcp", info.LocalString())
		if err != nil {
			log.Fatalln(fmt.Printf("Dial INTO local service error: %s", err))
		}

		client, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		handleClient(client, local)
	}
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
