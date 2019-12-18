package rpc

import (
	"Rabbit-OJ-Backend/utils/files"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/rpc"
)

func Register() {
	fmt.Println("[RPC] registering case service")

	if err := rpc.RegisterName("CaseService", new(CaseService)); err != nil {
		panic(err)
	}

	serverCertFile, err := files.CertFilePath("server.pem")
	if err != nil {
		panic(err)
	}
	serverKeyFile, err := files.CertFilePath("server.key")
	if err != nil {
		panic(err)
	}
	clientCertFile, err := files.CertFilePath("client.pem")
	if err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		return
	}

	certBytes, err := ioutil.ReadFile(clientCertFile)
	if err != nil {
		panic(err)
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}

	listener, err := tls.Listen("tcp", ":8090", config)
	if err != nil {
		panic(err)
	}

	rpc.Accept(listener)
}
