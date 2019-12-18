package initialize

import (
	"Rabbit-OJ-Backend/services/config"
	"Rabbit-OJ-Backend/utils/files"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

func Cert(role string) {
	certFile, err := files.CertFilePath(fmt.Sprintf("%s.pem", role))
	if err != nil {
		panic(err)
	}
	keyFile, err := files.CertFilePath(fmt.Sprintf("%s.key", role))
	if err != nil {
		panic(err)
	}
	clientCertFile, err := files.CertFilePath("client.pem")
	if err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
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

	config.X509KeyPair = cert
	config.ClientCertPool = clientCertPool
}
