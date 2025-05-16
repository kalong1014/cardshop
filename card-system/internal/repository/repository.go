package repository

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

// UserRepository 定义用户仓库接口
type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

// MerchantRepository 定义商户仓库接口
type MerchantRepository interface {
	FindByID(id uint) (*model.Merchant, error)
	FindByUserID(userID uint) (*model.Merchant, error)
	Create(merchant *model.Merchant) error
	Update(merchant *model.Merchant) error
	GetAll() ([]model.Merchant, error)
}

// PageRepository 定义页面仓库接口
type PageRepository interface {
	GetByID(id uint) (*model.Page, error)
	Create(page *model.Page) error
	Update(page *model.Page) error
	Delete(id uint) error
	GetByMerchantID(merchantID uint) ([]model.Page, error)
}

// UserRepositoryImpl 用户仓库实现
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// MerchantRepositoryImpl 商户仓库实现
type MerchantRepositoryImpl struct {
	db *gorm.DB
}

// NewMerchantRepository 创建商户仓库实例
func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &MerchantRepositoryImpl{db: db}
}

// PageRepositoryImpl 页面仓库实现
type PageRepositoryImpl struct {
	db *gorm.DB
}

// NewPageRepository 创建页面仓库实例
func NewPageRepository(db *gorm.DB) PageRepository {
	return &PageRepositoryImpl{db: db}
}
