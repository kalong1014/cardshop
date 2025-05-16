package repository

import (
	"card-system/internal/model"
)

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
