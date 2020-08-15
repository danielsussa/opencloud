package endpoint

import (
	"github.com/danielsussa/opencloud/shared"
	"golang.org/x/crypto/ssh"
)

func loadPresets(ch ssh.Channel, reqChans <-chan *ssh.Request) {
	defer ch.Close()
	select {
	case request := <- reqChans:
		if request == nil {
			return
		}
		ch.SendRequest(shared.LOAD_PRESETS_RESULT, false, []byte("presets result"))
	}
}
