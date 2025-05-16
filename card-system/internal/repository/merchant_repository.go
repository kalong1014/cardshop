package repository

import (
	"card-system/internal/model"
)

func (r *MerchantRepositoryImpl) FindByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := r.db.Preload("User").First(&merchant, id).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepositoryImpl) FindByUserID(userID uint) (*model.Merchant, error) {
	var merchant model.Merchant
	if err := r.db.Preload("User").Where("user_id = ?", userID).First(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepositoryImpl) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *MerchantRepositoryImpl) Update(merchant *model.Merchant) error {
	return r.db.Save(merchant).Error
}

func (r *MerchantRepositoryImpl) GetAll() ([]model.Merchant, error) {
	var merchants []model.Merchant
	if err := r.db.Preload("User").Find(&merchants).Error; err != nil {
		return nil, err
	}
	return merchants, nil
}
