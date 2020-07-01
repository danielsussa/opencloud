package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"log"
	"math/big"
	"time"
)

type RsaKeyPair struct {
	Private []byte
	Public  []byte
	Cert    []byte
}

type GerenateRsaKeyPair func(bitSize int) *RsaKeyPair

func rsaKeyGen(bitSize int) *RsaKeyPair {
	reader := rand.Reader
	key, err := rsa.GenerateKey(reader, bitSize)
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

	asn1Bytes, _ := asn1.Marshal(&key.PublicKey)

	var pubKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	return &RsaKeyPair{
		Cert:    pem.EncodeToMemory(certKey),
		Public:  pem.EncodeToMemory(pubKey),
		Private: pem.EncodeToMemory(privateKey),
	}
}
