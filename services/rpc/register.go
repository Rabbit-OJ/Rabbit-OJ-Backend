package rpc

import (
	"Rabbit-OJ-Backend/services/config"
	"crypto/tls"
	"fmt"
	"net/rpc"
)

func Register() {
	fmt.Println("[RPC] registering case service")

	if err := rpc.RegisterName("CaseService", new(CaseService)); err != nil {
		panic(err)
	}
	listener, err := tls.Listen("tcp", ":8090", &tls.Config{
		Certificates: []tls.Certificate{config.X509KeyPair},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    config.ClientCertPool,
	})
	if err != nil {
		panic(err)
	}

	rpc.Accept(listener)
}
