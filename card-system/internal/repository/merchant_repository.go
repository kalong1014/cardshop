package repository

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

// MerchantRepository 定义商户仓库接口
type MerchantRepository interface {
	Create(merchant *model.Merchant) error
	FindByID(id uint) (*model.Merchant, error)
	FindByUserID(userID uint) (*model.Merchant, error)
	Update(merchant *model.Merchant) error
	Delete(id uint) error
}

// merchantRepository 实现 MerchantRepository 接口
type merchantRepository struct {
	db *gorm.DB
}

// NewMerchantRepository 创建商户仓库实例
func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

// Create 创建商户
func (r *merchantRepository) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

// FindByID 根据ID查找商户
func (r *merchantRepository) FindByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.Preload("User").First(&merchant, id).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

// FindByUserID 根据用户ID查找商户
func (r *merchantRepository) FindByUserID(userID uint) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.Preload("User").Where("user_id = ?", userID).First(&merchant).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

// Update 更新商户信息
func (r *merchantRepository) Update(merchant *model.Merchant) error {
	return r.db.Save(merchant).Error
}

// Delete 删除商户
func (r *merchantRepository) Delete(id uint) error {
	return r.db.Delete(&model.Merchant{}, id).Error
}
