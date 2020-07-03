package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"strconv"
	"strings"
)


func(openAgent *OpenAgent) setConnectionPort() {
	cli, err := sshClient(openAgent)
	if err != nil {
		log.Fatal(err)
	}
	session, stdout, err := sshSession(cli, "connect_agent")
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Signal(ssh.SIGKILL); err != nil {
		log.Fatal(err)
	}

	var outBuf bytes.Buffer
	io.Copy(&outBuf, stdout)

	res := strings.Split(outBuf.String()," ")
	if res[0] != "connect_agent" {
		log.Fatal("cannot connect agent")
	}
	i,_ :=strconv.Atoi(res[1])
	openAgent.Port = i
}

func(openAgent *OpenAgent) startAdminProxy() {
	cli, err := sshClient(openAgent)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := cli.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", openAgent.Port))
	if err != nil {
		log.Fatalln(fmt.Printf("Listen open port ON remote server error: %s", err))
	}
	defer listener.Close()
	for {
		remoteConn, err := listener.Accept()
		if err != nil {
			remoteConn.Close()
			log.Println(err)
			continue
		}
		data, err := bufio.NewReader(remoteConn).ReadString('\n')
		if err != nil {
			remoteConn.Close()
			log.Println(err)
			continue
		}
		cmd := strings.TrimSpace(data)
		err = openAgent.commandHandler(cmd).Execute(remoteConn)
		if err != nil {
			remoteConn.Close()
			log.Println(err)
		}
	}
}


func sshClient(openAgent *OpenAgent)(*ssh.Client,error){
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
	return serverConn, err
}

func sshSession(conn *ssh.Client, command string)(*ssh.Session,io.Reader,error){
	session, err := conn.NewSession()
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
