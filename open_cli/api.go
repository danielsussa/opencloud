package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func main() {
	// Create client config
	config := &ssh.ClientConfig{
		User: "teste",
		Auth: []ssh.AuthMethod{
			ssh.Password("my_password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to ssh server
	conn, err := ssh.Dial("tcp", "localhost:9999", config)
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer conn.Close()
	// Create a session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("unable to create session: ", err)
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	command := ``

	if err := session.Start(command); err != nil {
		log.Fatal(err)
	}

	if err := session.Signal(ssh.SIGKILL); err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stderr, stderr)
	io.Copy(os.Stdout, stdout)

	if err := session.Wait(); err != nil {
		log.Println(err)
	}
}
