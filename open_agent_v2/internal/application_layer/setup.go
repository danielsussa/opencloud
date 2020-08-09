package application_layer

import sshConnection "github.com/danielsussa/opencloud/open_agent_v2/internal/application_layer/internal/ssh"

func Setup()error{
	return sshConnection.ConectToServer()
}
