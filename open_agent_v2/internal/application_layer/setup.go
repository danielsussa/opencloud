package application_layer

import (
	"github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer/api"
	sshConnection "github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer/internal/ssh"
	"github.com/danielsussa/opencloud/shared"
)

func Start()error{
	err :=  sshConnection.ConectToServer()
	if err != nil {
		return err
	}

	conn := sshConnection.GetSshConnection()
	_,_, err = conn.SendRequest(shared.LOAD_PRESETS, false, []byte("my agent info"))
	if err != nil {
		return err
	}

	err = api.Start(sshConnection.GetSshRequests())
	if err != nil {
		return err
	}
	return nil
}
