package rpc

import (
	"Rabbit-OJ-Backend/services/config"
	"crypto/tls"
	"fmt"
	"net/rpc"
)

type AnyType interface{}

func DialInit() (*rpc.Client, error) {
	tlsClient, err := tls.Dial("tcp", config.Global.Judger.Rpc, &tls.Config{
		RootCAs:            config.ClientCertPool,
		Certificates:       []tls.Certificate{config.X509KeyPair},
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := rpc.NewClient(tlsClient)
	return client, nil
}

func DialCall(serviceName, functionName string, request AnyType, response AnyType) error {
	client, err := DialInit()
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		return err
	}
	return client.Call(serviceName+"."+functionName, request, response)
}
