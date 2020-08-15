package api

import (
	"fmt"
	"github.com/danielsussa/opencloud/open_server_v2/internal/domain_layer/domain"
	"github.com/danielsussa/opencloud/shared"
	"golang.org/x/crypto/ssh"
)

func messageHandler(request *ssh.Request)error{
	switch request.Type {
	// Client can send a ping request to Agent respond
	case shared.PING:
		cmd, err := domain.NewPingCommand(request.Payload)
		if err != nil {
			return err
		}
		_, err = domain.PingDomain(cmd)
		if err != nil {
			return err
		}
		return nil
	case shared.LOAD_PRESETS:
		fmt.Println("hello")
	default:
		return nil
	}
	return nil
}
