package sshConnection

import (
	"golang.org/x/crypto/ssh"
	"net"
)

var sshConn ssh.Conn
var chanReq <-chan *ssh.Request

func ConectToServer()error{
	var err error
	sshConfig := &ssh.ClientConfig{
		// SSH connection username
		User: "teste",
		//Auth: []ssh.AuthMethod{
		//	// put here your private key path
		//	publicKeyFile("rsa"),
		//},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	tcpConn, err := net.Dial("tcp", "localhost:2222")
	if err != nil {
		return err
	}

	sshConn, _ ,chanReq, err = ssh.NewClientConn(tcpConn, "localhost:2222", sshConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetSshConnection()ssh.Conn {
	return sshConn
}

func GetSshRequests()<-chan *ssh.Request {
	return chanReq
}
