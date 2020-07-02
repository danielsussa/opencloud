package main

import (
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
			apiServer.CommandList = append(apiServer.CommandList, s.Command())
			command, err := getCommand(s.Command())
			if err != nil {
				log.Println(err)
				return
			}
			err = command.Execute(apiServer, s)
			if err != nil {
				log.Println(err)
				return
			}
			// execute command

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
