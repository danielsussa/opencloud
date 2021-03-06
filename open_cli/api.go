package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	apiCommand "github.com/danielsussa/opencloud/open_cli/command"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var password = "custom_password"

func main() {
	conf := loadConfig()
	command := apiCommand.ReturnCommand()

	// Create client config
	config := &ssh.ClientConfig{
		User: "opencloud",
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to ssh server
	conn, err := ssh.Dial("tcp", conf.Host, config)
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

	if err := session.Start(command.Request()); err != nil {
		log.Fatal(err)
	}

	if err := session.Signal(ssh.SIGKILL); err != nil {
		log.Fatal(err)
	}

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer

	io.Copy(&outBuf, stdout)
	io.Copy(&errBuf, stderr)

	if err := session.Wait(); err != nil {
		log.Println(err)
	}
	fmt.Println(command.Response(strings.Split(outBuf.String(), " ")))
}

type fileConfig struct {
	Host string
}

func loadConfig() (fc fileConfig) {
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fc = makeConfig()
	}
	json.Unmarshal(file, &fc)
	return fc
}

func makeConfig() (fc fileConfig) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter host (ex: localhost:9090): ")
	host, _ := reader.ReadString('\n')
	fc.Host = strings.Trim(host, "\n")

	b, _ := json.Marshal(fc)
	err := ioutil.WriteFile("config/config.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fc
}
