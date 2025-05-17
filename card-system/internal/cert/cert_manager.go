package cert

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"time"
)

type CertManager struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewCertManager() (*CertManager, error) {
	// 生成私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &CertManager{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

func (cm *CertManager) GenerateCardCert(cardNumber string) ([]byte, error) {
	// 创建证书模板
	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			CommonName:   cardNumber,
			Organization: []string{"Card System"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(5, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	// 生成证书
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, cm.publicKey, cm.privateKey)
	if err != nil {
		return nil, err
	}

	// PEM编码
	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	})

	return certPEM.Bytes(), nil
}

func (cm *CertManager) VerifyCardCert(cardNumber string, certData []byte) (bool, error) {
	// 解码PEM证书
	block, _ := pem.Decode(certData)
	if block == nil {
		return false, errors.New("无效的证书数据")
	}

	// 解析证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return false, err
	}

	// 验证证书有效期
	if time.Now().Before(cert.NotBefore) || time.Now().After(cert.NotAfter) {
		return false, nil
	}

	// 验证证书中的卡号
	if cert.Subject.CommonName != cardNumber {
		return false, nil
	}

	// 验证证书签名
	err = cert.CheckSignatureFrom(cert)
	return err == nil, err
}
