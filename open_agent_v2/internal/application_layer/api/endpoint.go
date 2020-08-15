package api

import (
	"github.com/danielsussa/opencloud/shared"
	"golang.org/x/crypto/ssh"
)

func Start(chRequests <-chan *ssh.Request) error {
	for {
		select {
		case request := <- chRequests:
			if request == nil {
				continue
			}
			switch request.Type {
			case shared.LOAD_PRESETS:
				err := loadPresetsEndpoint(request.Payload)
				if err != nil {
					return err
				}
			}
		}
	}
}
