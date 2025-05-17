package model

import (
	"gorm.io/gorm"
	"time"
)

// 域名模型
type Domain struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	MerchantID  uint       `gorm:"not null" json:"merchant_id"`
	DomainName  string     `gorm:"uniqueIndex" json:"domain_name"` // 完整域名（如m1-abc.cardshop.com）
	CertStatus  string     `gorm:"default:applied" json:"cert_status"` // 证书状态（applied/expired）
	CertExpire  time.Time  `json:"cert_expire"`
	CreatedAt   time.Time  `json:"created_at"`
}

// 自动申请证书（伪代码，需集成ACME库）
func (d *Domain) ApplyCertificate() error {
	// 使用lego库申请证书
	// client := acme.NewClient(...)
	// cert, err := client.GetCertificate(d.DomainName)
	// d.CertExpire = cert.Expiry
	// return err
	return nil
}