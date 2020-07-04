package command

import (
	"log"
	"net"
)

type pingCommand struct {

}

func (p pingCommand)Execute(conn net.Conn)error{
	log.Println("sucessfull receive ping message from server")
	_, err := conn.Write([]byte("pong\n"))
	if err != nil {
		return err
	}
	return nil
}
