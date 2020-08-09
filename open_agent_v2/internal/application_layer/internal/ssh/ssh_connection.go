package sshConnection

import "golang.org/x/crypto/ssh"

var sshConn *ssh.Client

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
	sshConn, err = ssh.Dial("tcp", "localhost:2222", sshConfig)
	if err != nil {
		return err
	}
	return nil
}

func GetSshConnection()*ssh.Client {
	return sshConn
}
