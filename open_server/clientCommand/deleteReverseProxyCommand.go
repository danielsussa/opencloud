package clientCommand

import (
	"errors"
	"fmt"
	"github.com/danielsussa/opencloud/open_server/sessionInfo"
	"github.com/danielsussa/opencloud/shared"
	"github.com/phayes/freeport"
	"strconv"
	"strings"
)

type deleteReverseProxyCommand struct {
	strArr []string
}

// delete_reverse_proxy agentName commandName
func (cmd deleteReverseProxyCommand) Execute() (string, error) {

}

func (cmd deleteReverseProxyCommand) Kind() string {
	return shared.DELETE_REVERSE_PROXY
}
