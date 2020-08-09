package domain

import (
	"encoding/json"
	"github.com/danielsussa/opencloud/open_server_v2/internal/domain_layer/ipush"
)

type PingCommandRequest struct {
	Agent string
}


func NewPingCommand(data []byte)(*PingCommandRequest, error){
	command := new(PingCommandRequest)
	err := json.Unmarshal(data, command)
	if err != nil {
		return nil, err
	}
	return command, nil
}


type PingCommandResponse struct {
	Status int
}

func PingDomain(cmd *PingCommandRequest)(*PingCommandResponse, error){
	err := ipush.Get().SendPingCommand(cmd.Agent)
	if err != nil {
		return nil, err
	}
	return &PingCommandResponse{Status: 200}, nil
}
