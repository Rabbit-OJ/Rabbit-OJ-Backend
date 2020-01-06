package config

import (
	"crypto/tls"
	"crypto/x509"
)

var (
	ClientCertPool *x509.CertPool
	X509KeyPair    tls.Certificate
)
