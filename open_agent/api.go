package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

type Config struct {
	bitSize       int
	User          string
	SshServerHost string
	SshServerPort int
}

func (c Config) ParseSshHostPort() string {
	return fmt.Sprintf("%s:%d", c.SshServerHost, c.SshServerPort)
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

	config := Config{
		bitSize:       2048,
		SshServerHost: "127.0.0.1",
		SshServerPort: 2223,
	}

	client := OpenAgent{
		CommandConnection: commandConnection,
		ApiCommandHandler: commandHandler,

		SshConnection:      sshConnection,
		AddSshReverseProxy: addReverseProxy,
	}
	client.Start(config)

}
