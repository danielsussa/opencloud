package command

import (
	sshUtils "github.com/danielsussa/opencloud/open_agent/ssh"
	"github.com/danielsussa/opencloud/shared"
	"log"
	"strconv"
)

type connectAgentCommand struct {
	cmd []string
}

func (p connectAgentCommand) Kind() string {
	return shared.CONNECT_AGENT
}

func (p connectAgentCommand)Execute()(string,error){
	port, err := strconv.Atoi(p.cmd[1])
	if err != nil {
		log.Fatal(err)
	}
	sshUtils.SetupSshPort(port)
	log.Println("sucessfull setting up port")
	return "pong", nil
}
