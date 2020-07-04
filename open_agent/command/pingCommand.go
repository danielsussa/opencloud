package command

import (
	"github.com/danielsussa/opencloud/shared"
	"log"
)

type pingCommand struct {

}

func (p pingCommand) Kind() string {
	return shared.PING
}

func (p pingCommand)Execute()(string,error){
	log.Println("sucessfull receive ping message from server")
	return "pong", nil
}
