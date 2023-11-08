package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"time"

	"dynamic-admission-webhook-example/router"
)

var (
	certFile string
	keyFile  string
	port     int
)

// Config contains the server (the webhook) cert and key.
type Config struct {
	CertFile string
	KeyFile  string
}

func init() {
	flag.StringVar(&certFile, "tls-cert-file", "", "File containing the default x509 Certificate for HTTPS. (CA cert, if any, concatenated after server cert).")
	flag.StringVar(&keyFile, "tls-private-key-file", "", "File containing the default x509 private key matching --tls-cert-file.")
	flag.IntVar(&port, "port", 443, "Secure port that the webhook listens on")
}

func main() {
	config := Config{
		CertFile: certFile,
		KeyFile:  keyFile,
	}
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router.InitRouter(),
		TLSConfig:      configTLS(config),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		return
	}
}

func configTLS(config Config) *tls.Config {
	sCert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{sCert},
		// TODO: uses mutual tls after we agree on what cert the apiserver should use.
		// ClientAuth:   tls.RequireAndVerifyClientCert,
	}
}
