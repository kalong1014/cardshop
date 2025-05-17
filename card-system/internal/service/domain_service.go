package service

import (
	"card-system/internal/model"
	"card-system/internal/repository"
	"fmt"
	"os"
	"time"
)

type DomainService struct {
	repo *repository.DomainRepo
}

func NewDomainService(repo *repository.DomainRepo) *DomainService {
	return &DomainService{repo: repo}
}

// 分配二级域名
func (s *DomainService) AllocateDomain(merchantID uint) (string, error) {
	domainName := fmt.Sprintf("m%d-%s.cardshop.com", merchantID, randString(6))
	domain := &model.Domain{
		MerchantID: merchantID,
		DomainName: domainName,
		CertStatus: model.CertStatusApplied,
		CertExpire: time.Now().Add(90 * 24 * time.Hour), // 证书默认有效期90天
	}
	if err := s.repo.Create(domain); err != nil {
		return "", err
	}
	go domain.ApplyCertificate() // 异步申请证书
	return domainName, nil
}

// 证书续期任务（定时执行）
func (s *DomainService) RenewCertificates() {
	// 每天凌晨1点检查过期证书
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			if time.Now().Hour() != 1 {
				continue
			}
			var domains []model.Domain
			utils.DB.Where("cert_status = ? AND cert_expire < ?",
				model.CertStatusApplied, time.Now().Add(7*24*time.Hour)).Find(&domains)
			for _, d := range domains {
				d.ApplyCertificate()
			}
		}
	}()
}

func (d *Domain) ApplyCertificate() error {
	certPEM, keyPEM, err := utils.ApplyCertificate(d.DomainName)
	if err != nil {
		d.CertStatus = model.CertStatusFailed
		return utils.DB.Save(d).Error
	}
	// 保存证书到服务器或数据库（示例：写入文件）
	os.WriteFile(fmt.Sprintf("/etc/letsencrypt/live/%s/fullchain.pem", d.DomainName), []byte(certPEM), 0644)
	os.WriteFile(fmt.Sprintf("/etc/letsencrypt/live/%s/privkey.pem", d.DomainName), []byte(keyPEM), 0600)

	d.CertStatus = model.CertStatusApplied
	d.CertExpire = time.Now().Add(90 * 24 * time.Hour)
	return utils.DB.Save(d).Error
}
