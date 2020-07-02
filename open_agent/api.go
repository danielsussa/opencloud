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
	sshConn    *ssh.Client
	Config     Config

	CommandConnection CommandConnection
	ApiCommandHandler ApiCommandHandler

	SshConnection      SshConnection
	AddSshReverseProxy AddSshReverseProxy
}

func (c OpenAgent) Start(config Config) {
	c.Config = config
	c.getOrGenerateRsaKeyGen()
	c.sshConn = c.SshConnection(config, c.rsaKeyPair)
	select {}
}

func main() {

	host := "127.0.0.1:2223"

	config := Config{
		bitSize:       2048,
		SshServerHost: host,
	}

	client := OpenAgent{
		CommandConnection: commandConnection,
		ApiCommandHandler: commandHandler,

		SshConnection:      sshConnection,
		AddSshReverseProxy: addReverseProxy,
	}
	client.Start(config)

}
