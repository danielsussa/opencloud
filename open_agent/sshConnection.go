package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)


func(openAgent OpenAgent) sshConnection() {
	key, err := ssh.ParsePrivateKey(openAgent.rsaKeyPair.Private)
	if err != nil {
		log.Fatalln(err)
	}

	sshConfig := &ssh.ClientConfig{
		// SSH connection username
		User:            openAgent.Config.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(key)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to SSH remote server using serverEndpoint
	serverConn, err := ssh.Dial("tcp", openAgent.Config.SshServerHost, sshConfig)
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}
	openAgent.sshClient = serverConn
	session, err := serverConn.NewSession()
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}
	if err := session.Start("init"); err != nil {
		log.Fatal(err)
	}

	openAgent.sshSession = session
}
