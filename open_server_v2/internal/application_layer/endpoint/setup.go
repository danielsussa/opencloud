package endpoint

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/danielsussa/opencloud/open_server_v2/internal/application_layer/internal"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func Start(){
	config := &ssh.ServerConfig{
		//PublicKeyCallback: publicKeyAuth,
		NoClientAuth: true,
	}
	config.AddHostKey(createSigner())


	// Once a ServerConfig has been configured, connections can be accepted.
	listener, err := net.Listen("tcp", "0.0.0.0:2222")
	if err != nil {
		log.Fatal(err)
	}
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			continue
		}
		go serveConnection(tcpConn, config)
	}
}

func serveConnection(conn net.Conn,config *ssh.ServerConfig)error{
	defer conn.Close()
	// Before use, a handshake must be performed on the incoming net.Conn.
	serverConn, _, requestsCh, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Printf("Failed to handshake (%s)", err)
		return err
	}
	internal.AddConnection("teste", serverConn)

	return endpoint(requestsCh)
}

func createSigner()ssh.Signer{
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.NewSignerFromKey(key)
	if err != nil {
		log.Fatal(err)
	}
	return signer
}

func publicKeyAuth(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error){
	return nil,nil
}
