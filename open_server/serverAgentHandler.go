package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"log"
)

func (apiServer *ApiServer) serverAgentHandler(config Config, errChan chan error) {
	forwardHandler := &ssh.ForwardedTCPHandler{}
	server := ssh.Server{
		Addr: config.ServerAgentPort,
		Handler: ssh.Handler(func(s ssh.Session) {
			// get command
			apiServer.CommandList = append(apiServer.CommandList, s.Command())
			command, err := getAgentCommand(s.Command())
			if err != nil {
				log.Println(err)
				return
			}
			err = command.Execute(apiServer, s)
			if err != nil {
				io.WriteString(s, fmt.Sprintf("%s\n", err.Error()))
				log.Println(err)
				return
			}
			// execute command

			sign := make(chan ssh.Signal)
			s.Signals(sign)
			select {
			case k := <-sign:
				if k == ssh.SIGKILL {
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
		RequestHandlers: map[string]ssh.RequestHandler{
			"tcpip-forward":        forwardHandler.HandleSSHRequest,
			"cancel-tcpip-forward": forwardHandler.HandleSSHRequest,
		},
	}
	go func() {
		errChan <- server.ListenAndServe()
	}()
}
