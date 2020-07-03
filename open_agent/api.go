package main

type Config struct {
	bitSize       int
	User          string
	SshServerHost string
}


type OpenAgent struct {
	rsaKeyPair *RsaKeyPair
	Config Config

	Port int
}

func (c OpenAgent) Start() {
	c.getOrGenerateRsaKeyGen()
	c.setConnectionPort()
	c.startAdminProxy()
}

func main() {

	confFile := loadConfig()

	config := Config{
		bitSize:       2048,
		SshServerHost: confFile.Host,
		User: confFile.AgentName,
	}

	client := OpenAgent{
		Config: config,
	}
	client.Start()

}

