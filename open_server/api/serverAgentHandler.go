package api

import (
	"encoding/base64"
	"fmt"
	errorhandler "github.com/danielsussa/opencloud/open_server/error"
	"io"
	"log"

	"github.com/danielsussa/opencloud/open_server/agentCommand"
	"github.com/danielsussa/opencloud/open_server/data"
	"github.com/gliderlabs/ssh"
)

func (apiServer *ApiServer) serverAgentHandler(config Config) {
	forwardHandler := &ssh.ForwardedTCPHandler{}
	server := ssh.Server{
		Addr: config.ServerAgentPort,
		Handler: ssh.Handler(func(s ssh.Session) {
			// get command
			command, err := agentCommand.GetAgentCommand(s.Command())
			if err != nil {
				log.Println(err)
				return
			}
			_, err = command.Execute(s)
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
			return data.AnyUserHasKey(keyEncoded)
		}),
		RequestHandlers: map[string]ssh.RequestHandler{
			"tcpip-forward":        forwardHandler.HandleSSHRequest,
			"cancel-tcpip-forward": forwardHandler.HandleSSHRequest,
		},
	}
	go func() {
		errorhandler.GetErrChan() <- server.ListenAndServe()
	}()
}
