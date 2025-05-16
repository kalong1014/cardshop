// internal/repository/merchant_repository.go
package repository

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) FindByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := r.db.Preload("User").First(&merchant, id).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) FindByUserID(userID uint) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := r.db.Preload("User").Where("user_id = ?", userID).First(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *MerchantRepository) Update(merchant *model.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *MerchantRepository) GetAll() ([]model.Merchant, error) {
	var merchants []model.Merchant
	if err := r.db.Preload("User").Find(&merchants).Error; err != nil {
		return nil, err
	}
	return merchants, nil
}
