package main

import (
	"fmt"
	"github.com/danielsussa/opencloud/open_server/clientCommand"
	"github.com/danielsussa/opencloud/shared"
	"github.com/gliderlabs/ssh"
	"io"
	"log"
)

func (apiServer *ApiServer) serverClientHandler(config Config, errChan chan error) {
	server := ssh.Server{
		Addr: config.ServerClientPort,
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			if password != apiServer.Config.ServerClientPassword {
				return false
			}
			return true
		},
		Handler: ssh.Handler(func(s ssh.Session) {
			// get command
			command, err := clientCommand.GetClientCommand(s.Command())
			if err != nil {
				io.WriteString(s, fmt.Sprintf("%s 400 %s\n", shared.NIL, err.Error()))
				log.Println(err)
				return
			}

			// execute command
			res, err := command.Execute()
			if err != nil {
				io.WriteString(s, fmt.Sprintf("%s 500 %s\n", command.Kind(), err.Error()))
				log.Println(err)
				return
			}
			io.WriteString(s, fmt.Sprintf("%s 200 %s\n", command.Kind(), res))

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
