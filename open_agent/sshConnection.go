package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/danielsussa/opencloud/open_agent/command"
	sshUtils "github.com/danielsussa/opencloud/open_agent/ssh"
	"github.com/danielsussa/opencloud/shared"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"strings"
)

func (openAgent *OpenAgent) setConnectionPort() {
	keyPair := sshUtils.GetOrGenerateRsaKeyGen()
	cli, err := sshUtils.SshClient(openAgent.Config.SshServerHost, openAgent.Config.User, keyPair.Private)
	if err != nil {
		if strings.Contains(err.Error(), "handshake failed"){
			sshUtils.ConsoleMessage(keyPair)
		}
		log.Fatal(err)
	}
	session, stdout, err := sshUtils.SshSession(cli, shared.CONNECT_AGENT)
	if err != nil {
		log.Fatal(err)
	}

	if err := session.Signal(ssh.SIGKILL); err != nil {
		log.Fatal(err)
	}


	var outBuf bytes.Buffer
	io.Copy(&outBuf, stdout)

	commandLines := strings.Split(outBuf.String(), "\n")

	for _, commandLine := range commandLines{
		if commandLine == ""{
			continue
		}
		cmd := command.CommandHandler(strings.TrimSpace(commandLine))
		_, err := cmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (openAgent *OpenAgent) startAdminProxy() {
	keyPair := sshUtils.GetOrGenerateRsaKeyGen()
	cli, err := sshUtils.SshClient(openAgent.Config.SshServerHost, openAgent.Config.User, keyPair.Private)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := cli.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", sshUtils.GetSshPort()))
	if err != nil {
		log.Fatalln(fmt.Printf("Listen open port ON remote server error: %s", err))
	}
	defer listener.Close()
	for {
		remoteConn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		data, err := bufio.NewReader(remoteConn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmd := command.CommandHandler(strings.TrimSpace(data))
		msg, err := cmd.Execute()
		if err != nil {
			remoteConn.Write([]byte(fmt.Sprintf("%s 500 %s\n",cmd.Kind(), err.Error())))
			continue
		}
		remoteConn.Write([]byte(fmt.Sprintf("%s 200 %s\n",cmd.Kind(), msg)))
	}
}
