package rpc

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/utils/files"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/rpc"
)

var (
	ClientCert     tls.Certificate
	ClientCertPool *x509.CertPool
)

type AnyType interface{}

func CertInit() {
	clientCertFile, err := files.CertFilePath("client.pem")
	if err != nil {
		panic(err)
	}
	clientKeyFile, err := files.CertFilePath("client.key")
	if err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		panic(err)
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

	ClientCert, ClientCertPool = cert, clientCertPool
}

func DialInit() (*rpc.Client, error) {
	if ClientCertPool == nil {
		CertInit()
	}

	conf := &tls.Config{
		RootCAs:            ClientCertPool,
		Certificates:       []tls.Certificate{ClientCert},
		InsecureSkipVerify: true,
	}

	tlsClient, err := tls.Dial("tcp", config.Global.Rpc, conf)
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
