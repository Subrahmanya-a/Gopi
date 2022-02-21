package config

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
)
/*
	This function injects required certificates(PEM)
	
	:rtype *http.Client
	:return client : HTTP Client having required certificat details for request authorization
	*/

	// Get required certs from respective files : .pem

var (
	certFile = flag.String("certFile", "../certs/api.pem", "A PEM eoncoded certificate file.")
)
//path absolute
func CertsInjectionsPem() *http.Client {
	flag.Parse()
	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *certFile)
	if err != nil {
		fmt.Println(err)
	}
	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}
	return client
}
