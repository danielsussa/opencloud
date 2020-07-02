package main

import (
	"github.com/gliderlabs/ssh"
	"io"
)

func (apiServer *ApiServer) serverAgentHandler(config Config, errChan chan error) {
	server := ssh.Server{
		Addr: config.ServerClientPort,
		Handler: ssh.Handler(func(s ssh.Session) {
			io.WriteString(s, "Remote forwarding available...\n")
			agentSession := apiServer.agentSession[s.PublicKey()]
			agentSession.Session = &s
			sign := make(chan ssh.Signal)
			s.Signals(sign)
			select {
			case k := <-sign:
				if k == ssh.SIGKILL {
					return
				}
			}
		}),
		PublicKeyHandler: ssh.PublicKeyHandler(func(ctx ssh.Context, key ssh.PublicKey) bool {
			if _, ok := apiServer.agentSession[key]; !ok {
				return false
			}
			return true
		}),
	}
	go func() {
		errChan <- server.ListenAndServe()
	}()
}
