// internal/repository/page_repository.go
package repository

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

type PageRepository interface {
	GetByID(id uint) (*model.Page, error)
	Create(page *model.Page) error
	Update(page *model.Page) error
	Delete(id uint) error
	GetByMerchantID(merchantID uint) ([]model.Page, error)
}

type PageRepositoryImpl struct {
	db *gorm.DB
}

func NewPageRepository(db *gorm.DB) PageRepository {
	return &PageRepositoryImpl{db: db}
}

func (r *PageRepositoryImpl) GetByID(id uint) (*model.Page, error) {
	var page model.Page
	if err := r.db.First(&page, id).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *PageRepositoryImpl) Create(page *model.Page) error {
	return r.db.Create(page).Error
}

func (r *PageRepositoryImpl) Update(page *model.Page) error {
	return r.db.Save(page).Error
}

func (r *PageRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&model.Page{}, id).Error
}

func (r *PageRepositoryImpl) GetByMerchantID(merchantID uint) ([]model.Page, error) {
	var pages []model.Page
	if err := r.db.Where("merchant_id = ?", merchantID).Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}
