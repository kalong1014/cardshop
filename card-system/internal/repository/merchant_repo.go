package repository

import (
	"card-system/internal/model"
	"gorm.io/gorm"
)

type MerchantRepo struct {
	db *gorm.DB
}

func NewMerchantRepo(db *gorm.DB) *MerchantRepo {
	return &MerchantRepo{db: db}
}

// 创建商户
func (r *MerchantRepo) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

// 根据用户ID查询商户
func (r *MerchantRepo) GetByUserID(userID uint) (*model.Merchant, error) {
	var m model.Merchant
	err := r.db.Preload("User").Where("user_id = ?", userID).First(&m).Error
	return &m, err
}

// 更新商户状态
func (r *MerchantRepo) UpdateStatus(merchantID uint, status model.MerchantStatus) error {
	return r.db.Model(&model.Merchant{}).Where("id = ?", merchantID).Update("status", status).Error
}