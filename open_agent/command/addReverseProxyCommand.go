package command

import (
	"fmt"
	"github.com/danielsussa/opencloud/open_agent/reverseProxy"
	sshUtils "github.com/danielsussa/opencloud/open_agent/ssh"
	"github.com/danielsussa/opencloud/shared"
	"log"
	"strconv"
)

type addReverseProxyCommand struct {
	cmd []string
}


func (p addReverseProxyCommand) Kind() string {
	return shared.ADD_REVERSE_PROXY
}

// add_reverse_proxy agent proxyName 8080 52738
func (p addReverseProxyCommand) Execute() (string, error) {
	// read message
	user := p.cmd[1]
	proxyName := p.cmd[2]
	localPort , err := strconv.Atoi(p.cmd[3])
	if err != nil {
		return "",err
	}
	remotePort , err := strconv.Atoi(p.cmd[4])
	if err != nil {
		return "",err
	}
	keyPair := sshUtils.GetOrGenerateRsaKeyGen()
	remoteHost := fmt.Sprintf("localhost:%d", remotePort)
	localHost := fmt.Sprintf("localhost:%d", localPort)

	err = reverseProxy.NewReverseProxy(remoteHost, localHost, user,proxyName, keyPair.Private)
	if err != nil {
		return "",err
	}
	log.Println(fmt.Sprintf("Sucessfull setting reverse proxy local:%d remote:%d", localPort, remotePort))
	return "",nil
}
