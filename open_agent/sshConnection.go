package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

type SshConnection func(config Config, keyPair *RsaKeyPair) *ssh.Client

func sshConnection(config Config, keyPair *RsaKeyPair) *ssh.Client {
	key, err := ssh.ParsePrivateKey(keyPair.Private)
	if err != nil {
		log.Fatalln(err)
	}

	sshConfig := &ssh.ClientConfig{
		// SSH connection username
		User:            config.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(key)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to SSH remote server using serverEndpoint
	serverConn, err := ssh.Dial("tcp", config.ParseSshHostPort(), sshConfig)
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}
	return serverConn
}
