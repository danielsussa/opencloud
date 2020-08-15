package endpoint

import (
	"fmt"
	"github.com/danielsussa/opencloud/shared"
	"golang.org/x/crypto/ssh"
)

func endpoint(chRequests <-chan *ssh.Request) error {
	for {
		select {
		case request := <- chRequests:
			if request == nil {
				continue
			}
			switch request.Type {
			case shared.LOAD_PRESETS:
				fmt.Println("hello")
			}
		}
	}
}

