package api

import (
	"encoding/json"
	"fmt"
	"github.com/danielsussa/opencloud/open_server_v2/internal/domain_layer/domain"
	"github.com/danielsussa/opencloud/open_server_v2/internal/domain_layer/domain/commands"
	"github.com/danielsussa/opencloud/shared"
	"golang.org/x/crypto/ssh"
)

func messageHandler(ch ssh.Channel, request *ssh.Request)error{
	switch request.Type {
	// Client can send a ping request to Agent respond
	case shared.PING:
		cmd, err := domain.NewPingCommand(request.Payload)
		if err != nil {
			return failureResponse(err, ch, request)
		}
		res,err := domain.PingDomain(cmd)
		if err != nil {
			return failureResponse(err, ch, request)
		}
		return successResponse(res, ch, request)
	case shared.LOAD_PRESETS:
		fmt.Println("hello")
	default:
		return nil
	}
	return nil
}

func successResponse(response interface{}, ch ssh.Channel, request *ssh.Request)error{
	b, _ := json.Marshal(response)
	_, err := ch.SendRequest(request.Type, false, b)
	if err != nil {
		return err
	}
	return nil
}

func failureResponse(currError error, ch ssh.Channel, request *ssh.Request)error{
	errResponse := commands.ErrorResponse{
		Status: 400,
		Reason: currError.Error(),
	}
	b, _ := json.Marshal(errResponse)
	_, err := ch.SendRequest(request.Type, false, b)
	if err != nil {
		return err
	}
	return nil
}
