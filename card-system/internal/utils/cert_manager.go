package utils

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

// 申请证书
func ApplyCertificate(domain string) (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	config := lego.Config{
		Email:      "admin@cardshop.com",
		PrivateKey: privateKey,
		HTTPClient: &http.Client{},
	}
	client, err := lego.NewClient(config)
	if err != nil {
		return "", "", err
	}

	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return "", "", err
	}
	client = client.WithRegistration(reg)

	certRequest := lego.CertificateRequest{
		Domains: []string{domain},
		Subject: &pkix.Name{
			Organization: []string{"CardShop"},
		},
	}
	certificates, err := client.Certificate.Get(certRequest)
	if err != nil {
		return "", "", err
	}

	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificates.Certificate}),
		string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})),
		nil
}