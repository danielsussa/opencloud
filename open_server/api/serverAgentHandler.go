package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
)

func (apiServer *ApiServer) serverAgentHandler(config Config, errChan chan error) {
	server := ssh.Server{
		Addr: config.ServerAgentPort,
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			return true
		},
		Handler: ssh.Handler(func(s ssh.Session) {
			fmt.Println(s.Command())
			io.WriteString(s, "writing data back...\n")
			sign := make(chan ssh.Signal)
			s.Signals(sign)
			select {
			case k := <-sign:
				if k == ssh.SIGKILL {
					return
				}
			}
		}),
	}

	go func() {
		errChan <- server.ListenAndServe()
	}()
}
