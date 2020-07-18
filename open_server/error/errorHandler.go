package errorhandler

import (
	"log"
)

var universalChan chan error

func init() {
	universalChan = make(chan error)
}

func GetErrChan() chan error {
	return universalChan
}

func ListenForErrors() {
	select {
	case err := <-universalChan:
		if err != nil {
			log.Fatal(err)
		}
	}
}
