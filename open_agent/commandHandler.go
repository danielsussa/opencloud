package main

import (
	"net"
	"strings"
)

type ApiCommand interface {
	Execute(conn net.Conn)error
}

type pingCommand struct {

}

func (p pingCommand)Execute(conn net.Conn)error{
	_, err := conn.Write([]byte("pong\n"))
	if err != nil {
		return err
	}
	return nil
}

func(op OpenAgent) commandHandler(cmd string) ApiCommand{
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[0] {
	case "ping":
		return pingCommand{}
	}
	return nil
}
