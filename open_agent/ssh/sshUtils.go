package sshUtils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
)

var sshClient *ssh.Client
var port int

func SetupSshPort(p int){
	if port != 0 {
		return
	}
	port = p
}

func GetSshPort()int{
	return port
}

func SshClient(host string,user string, pkey []byte) (*ssh.Client, error) {
	if sshClient != nil {
		return sshClient, nil
	}
	key, err := ssh.ParsePrivateKey(pkey)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		// SSH connection username
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(key)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to SSH remote server using serverEndpoint
	sshClient, err = ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, err
	}
	return sshClient, err
}

func SshSession(client *ssh.Client, command string) (*ssh.Session, io.Reader, error) {
	session, err := client.NewSession()
	if err != nil {
		log.Fatalln(fmt.Printf("Dial INTO remote server error: %s", err))
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Start(command); err != nil {
		log.Fatal(err)
	}
	return session, stdout, err
}
