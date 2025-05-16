package service

import (
	"card-system/internal/model"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type PageService struct {
	db *gorm.DB
}

func NewPageService(db *gorm.DB) *PageService {
	return &PageService{db: db}
}

// 创建页面
func (s *PageService) CreatePage(merchantID uint, name string) (*model.Page, error) {
	page := &model.Page{
		MerchantID: merchantID,
		Name:       name,
		Status:     "draft",
		Elements:   "[]",
	}

	result := s.db.Create(page)
	if result.Error != nil {
		return nil, result.Error
	}

	return page, nil
}

// 获取商户所有页面
func (s *PageService) GetMerchantPages(merchantID uint) ([]model.Page, error) {
	var pages []model.Page
	result := s.db.Where("merchant_id = ?", merchantID).Find(&pages)
	if result.Error != nil {
		return nil, result.Error
	}

	return pages, nil
}

// 获取单个页面
func (s *PageService) GetPageByID(id uint, merchantID uint) (*model.Page, error) {
	var page model.Page
	result := s.db.Where("id = ? AND merchant_id = ?", id, merchantID).First(&page)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("页面不存在")
		}
		return nil, result.Error
	}

	return &page, nil
}

// 保存页面
func (s *PageService) SavePage(id uint, merchantID uint, name string, elements interface{}) error {
	// 转换元素为JSON字符串
	elementsJSON, err := json.Marshal(elements)
	if err != nil {
		return err
	}

	// 更新页面
	result := s.db.Model(&model.Page{}).Where("id = ? AND merchant_id = ?", id, merchantID).Updates(map[string]interface{}{
		"name":     name,
		"elements": string(elementsJSON),
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("页面不存在或无权限访问")
	}

	return nil
}
