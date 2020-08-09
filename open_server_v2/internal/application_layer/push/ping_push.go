package push

import (
	"github.com/danielsussa/opencloud/open_server_v2/internal/application_layer/internal"
	"github.com/danielsussa/opencloud/shared"
)

func PingPush(agent string) error {
	ch := internal.GetChannelByAgent(agent)
	_, err := ch.SendRequest(shared.PING, false, nil)
	if err != nil {
		return err
	}
	return nil
}
