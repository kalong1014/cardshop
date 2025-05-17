package utils

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/lego"
)

func GenerateCertificate(domain string) (string, string, error) {
	// 配置更新为新版 lego 接口
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	config := lego.Config{
		Email:      "admin@example.com", // 保留必要字段
		PrivateKey: privateKey,
		// 移除过时的 Registration 字段
	}

	// 使用默认 DNS 挑战（示例）
	provider, err := dns01.NewDNSProvider("")
	if err != nil {
		return "", "", err
	}

	client, err := lego.NewClient(&config)
	if err != nil {
		return "", "", err
	}
	client.Challenge.Set(challenge.DNS01, provider)

	certRequest := lego.CertificateRequest{
		Domains: []string{domain},
		// 移除过时的 Host 字段
		PrivateKeyType: certcrypto.RSA2048,
	}

	certBytes, keyBytes, err := client.Certificate.Get(certRequest)
	if err != nil {
		return "", "", err
	}

	return string(certBytes), string(keyBytes), nil
}