package service

import (
	"card-system/model"
	"card-system/repository"
)

type TemplateService struct {
	repo *repository.TemplateRepo
}

func NewTemplateService(repo *repository.TemplateRepo) *TemplateService {
	return &TemplateService{repo: repo}
}

// 保存模板
func (s *TemplateService) SaveTemplate(merchantID uint, name string, data string) error {
	return s.repo.Create(&model.PageTemplate{
		MerchantID:  merchantID,
		Name:        name,
		LayoutData:  data,
		IsDefault:   false,
	})
}

// 设置默认模板
func (s *TemplateService) SetDefaultTemplate(templateID uint) error {
	// 先取消其他模板的默认状态
	if err := s.repo.Update(&model.PageTemplate{IsDefault: false}, 
		"merchant_id = ? AND is_default = ?", 
		model.GetMerchantIDFromContext(c), true); err != nil {
		return err
	}
	return s.repo.Update(&model.PageTemplate{IsDefault: true}, "id = ?", templateID)
}