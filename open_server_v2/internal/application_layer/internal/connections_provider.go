package internal

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

var connections map[string]*ssh.ServerConn

func init(){
	connections = make(map[string]*ssh.ServerConn)
}

func AddConnection(key string, conn *ssh.ServerConn){
	connections[key] = conn
}

func GetConnection(key string)(*ssh.ServerConn, error) {
	if conn, ok := connections[key]; ok {
		return conn, nil
	}
	return nil, fmt.Errorf("cannot find connection")
}
