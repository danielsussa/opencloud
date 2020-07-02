package main

import (
	"encoding/base64"
	"github.com/gliderlabs/ssh"
	"io"
	"log"
)

func (apiServer *ApiServer) serverAgentHandler(config Config, errChan chan error) {
	server := ssh.Server{
		Addr: config.ServerAgentPort,
		Handler: ssh.Handler(func(s ssh.Session) {
			io.WriteString(s, "Remote forwarding available...\n")
			keyEncoded := base64.StdEncoding.EncodeToString(s.PublicKey().Marshal())
			agentSession := apiServer.agentSession[keyEncoded]
			agentSession.Session = &s
			sign := make(chan ssh.Signal)
			s.Signals(sign)
			select {
			case k := <-sign:
				if k == ssh.SIGKILL {
					delete(apiServer.agentSession, keyEncoded)
					return
				}
			}
		}),
		ReversePortForwardingCallback: ssh.ReversePortForwardingCallback(func(ctx ssh.Context, host string, port uint32) bool {
			log.Println("attempt to bind", host, port, "granted")
			return true
		}),
		LocalPortForwardingCallback: ssh.LocalPortForwardingCallback(func(ctx ssh.Context, dhost string, dport uint32) bool {
			log.Println("Accepted forward", dhost, dport)
			return true
		}),
		PublicKeyHandler: ssh.PublicKeyHandler(func(ctx ssh.Context, key ssh.PublicKey) bool {
			keyEncoded := base64.StdEncoding.EncodeToString(key.Marshal())
			if _, ok := apiServer.agentSession[keyEncoded]; !ok {
				return false
			}
			return true
		}),
	}
	go func() {
		errChan <- server.ListenAndServe()
	}()
}
