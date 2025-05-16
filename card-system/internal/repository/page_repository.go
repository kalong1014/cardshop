package repository

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

type PageRepository struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{db: db}
}

// 创建页面
func (r *PageRepository) Create(page *model.Page) error {
	return r.db.Create(page).Error
}

// 获取商户所有页面
func (r *PageRepository) GetByMerchantID(merchantID uint) ([]model.Page, error) {
	var pages []model.Page
	err := r.db.Where("merchant_id = ?", merchantID).Find(&pages).Error
	return pages, err
}

// 获取单个页面
func (r *PageRepository) GetByID(id uint, merchantID uint) (*model.Page, error) {
	var page model.Page
	err := r.db.Where("id = ? AND merchant_id = ?", id, merchantID).First(&page).Error
	return &page, err
}

// 更新页面
func (r *PageRepository) Update(page *model.Page) error {
	return r.db.Save(page).Error
}
