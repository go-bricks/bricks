package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

type TLSConfigType string

const (
	TLSClientConfig TLSConfigType = "tlsClientConfig"
	TLSserverConfig TLSConfigType = "tlsserverConfig"
)

func LoadTLSCredentials(CaCertFile string, KeyFile string, CertFile string, tlsConfig TLSConfigType) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed client's certificate
	pemClientCA, err := ioutil.ReadFile(CaCertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// Load server's certificate and private key
	tlsCert, err := tls.LoadX509KeyPair(CertFile, KeyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{}
	switch tlsConfig {
	case TLSserverConfig:
		config = &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ClientCAs:    certPool,
		}
	default:
		config = &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			RootCAs:      certPool,
		}
	}

	return credentials.NewTLS(config), nil
}
