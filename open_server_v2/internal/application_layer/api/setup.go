package api

import (
	"crypto/rand"
	"crypto/rsa"
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
		// Before use, a handshake must be performed on the incoming net.Conn.
		_, chans, chanReq, err := ssh.NewServerConn(tcpConn, config)
		if err != nil {
			log.Printf("Failed to handshake (%s)", err)
			continue
		}
		go ssh.DiscardRequests(chanReq)
		go getOpenedChannel(chans, tcpConn)
	}
}

func getOpenedChannel(chans <-chan ssh.NewChannel, conn net.Conn) {
	select {
	case currChan := <- chans: // happens when client sshServerConn.OpenChannel
		if currChan == nil {
			conn.Close()
			return
		}
		ch, reqChan, err := currChan.Accept()
		if err != nil {
			log.Println(err)
		}

		for {
			select {
			case currReq := <- reqChan:
				if currReq == nil {
					conn.Close()
					return
				}
				err = messageHandler(ch, currReq)
				if err != nil {
					currReq.Reply(false, []byte("heeeelllo"))
				}
				currReq.Reply(true, []byte("heeeelllo"))
			}
		}
	}
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
