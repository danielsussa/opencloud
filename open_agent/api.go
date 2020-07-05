package main

type Config struct {
	bitSize       int
	User          string
	SshServerHost string
}


type OpenAgent struct {
	Config Config
}

func (c OpenAgent) Start() {
	c.setConnectionPort()
	c.startAdminProxy()
}

func main() {

	confFile := loadConfig()

	config := Config{
		SshServerHost: confFile.Host,
		User: confFile.AgentName,
	}

	client := OpenAgent{
		Config: config,
	}
	client.Start()

}

