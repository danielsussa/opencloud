package main

import (
	"golang.org/x/crypto/ssh"
)

type Config struct {
	bitSize       int
	User          string
	SshServerHost string
}


type OpenAgent struct {
	rsaKeyPair *RsaKeyPair

	sshSession *ssh.Session
	sshClient  *ssh.Client
	Config Config

	CommandConnection CommandConnection
	ApiCommandHandler ApiCommandHandler

	AddSshReverseProxy AddSshReverseProxy
}

func (c OpenAgent) Start(config Config) {
	c.Config = config
	c.getOrGenerateRsaKeyGen()
	c.sshConnection()
	select {}
}

func main() {

	confFile := loadConfig()

	config := Config{
		bitSize:       2048,
		SshServerHost: confFile.Host,
		User: confFile.AgentName,
	}

	client := OpenAgent{
		CommandConnection: commandConnection,
		ApiCommandHandler: commandHandler,
		AddSshReverseProxy: addReverseProxy,
	}
	client.Start(config)

}

