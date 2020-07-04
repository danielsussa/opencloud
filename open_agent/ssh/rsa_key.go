package sshUtils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"math/big"
	"time"
)

type RsaKeyPair struct {
	Private []byte
	Public  []byte
	RsaPublic []byte
	Cert    []byte
}

var keyPair *RsaKeyPair

func GetOrGenerateRsaKeyGen()*RsaKeyPair {
	if keyPair != nil {
		return keyPair
	}
	// try load file
	keyPair = loadKeyFile()
	if keyPair != nil {
		return keyPair
	}

	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	tml := x509.Certificate{
		// you can add any attr that you need
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(5, 0, 0),
		// you have to generate a different serial number each execution
		SerialNumber:          big.NewInt(123123),
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &key.PublicKey, key)
	if err != nil {
		log.Fatal("Certificate cannot be created.", err.Error())
	}

	var certKey = &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	}

	asn1Bytes, _ := asn1.Marshal(key.PublicKey)

	var pubKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	signer, err := ssh.ParsePrivateKey(pem.EncodeToMemory(privateKey))
	if err != nil {
		log.Fatalln(err)
	}

	keyPair =  &RsaKeyPair{
		Cert:    pem.EncodeToMemory(certKey),
		Public:  pem.EncodeToMemory(pubKey),
		RsaPublic: signer.PublicKey().Marshal(),
		Private: pem.EncodeToMemory(privateKey),
	}
	saveKeyFile(keyPair)
	return keyPair
}

func ConsoleMessage(keyPair *RsaKeyPair) {
	encodedPub := base64.StdEncoding.EncodeToString(keyPair.RsaPublic)
	log.Println(fmt.Sprintf("to add key on server use the command: opencli -command=new_agent -agent=agentName -key=%s", encodedPub))
}

func saveKeyFile(keyPair *RsaKeyPair) {
	b, _ := json.Marshal(keyPair)
	err := ioutil.WriteFile("config/rsaKeyPair.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func loadKeyFile() (rsa *RsaKeyPair) {
	file, err := ioutil.ReadFile("config/rsaKeyPair.json")
	if err != nil {
		return nil
	}
	err = json.Unmarshal(file, &rsa)
	return rsa
}
