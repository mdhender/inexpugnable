// inexpugnable - an esmtp server
// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2016-2019 GuerrillaMail.com.

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"os"
	"time"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "genkey",
		Short: "generate a certificate for local testing",
		Run:   genX509KeyPair,
	})
}

// genX509KeyPair cribbed from https://go.dev/src/crypto/tls/generate_cert.go
func genX509KeyPair(cmd *cobra.Command, args []string) {
	now := time.Now()
	template := &x509.Certificate{
		SerialNumber: big.NewInt(now.Unix()),
		Subject: pkix.Name{
			CommonName:         "inexpugnable.com",
			Country:            []string{"USA"},
			Organization:       []string{"inexpugnable.com"},
			OrganizationalUnit: []string{"mail-daemon"},
		},
		NotBefore:             now,
		NotAfter:              now.AddDate(0, 3, 1), // Valid for three months and a day
		SubjectKeyId:          []byte("mail.inexpugnable.com"),
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, priv.Public(), priv)
	if err != nil {
		log.Fatal(err)
	}

	certFile, keyFile := "pem/inexpugnable.com.pem", "pem/inexpugnable.com.key"
	if certOut, err := os.Create(certFile); err != nil {
		log.Fatalf("Failed to open %s for writing: %v", certFile, err)
	} else if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		log.Fatalf("Failed to write data to %s: %v", certFile, err)
	} else if err := certOut.Close(); err != nil {
		log.Fatalf("Error closing %s: %v", certFile, err)
	}
	log.Printf("wrote %s\n", certFile)

	if keyOut, err := os.OpenFile(keyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600); err != nil {
		log.Fatalf("Failed to open %s for writing: %v", keyFile, err)
	} else if privBytes, err := x509.MarshalPKCS8PrivateKey(priv); err != nil {
		log.Fatalf("Unable to marshal private key: %v", err)
	} else if err = pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		log.Fatalf("Failed to write data to %s: %v", keyFile, err)
	} else if err = keyOut.Close(); err != nil {
		log.Fatalf("Error closing %s: %v", keyFile, err)
	}
	log.Printf("wrote %s\n", keyFile)
}
