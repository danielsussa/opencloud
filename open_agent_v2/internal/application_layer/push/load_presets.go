package push

import (
	sshConnection "github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer/internal/ssh"
	"github.com/danielsussa/opencloud/shared"
)

func LoadPresets()error{
	sshConn := sshConnection.GetSshConnection()

	currCh, chanReq, err := sshConn.OpenChannel(shared.LOAD_PRESETS, []byte("do gato"))
	if err != nil {
		return err
	}
	//go ssh.DiscardRequests(chanReq)
	ok, err := currCh.SendRequest(shared.LOAD_PRESETS, true, []byte("do gato"))

	if err != nil && !ok {
		return err
	}

	currCh.Close()


	return nil
}
