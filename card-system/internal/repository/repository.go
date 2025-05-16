package repository

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

type MerchantRepository interface {
	FindByID(id uint) (*model.Merchant, error)
	FindByUserID(userID uint) (*model.Merchant, error)
	Create(merchant *model.Merchant) error
	Update(merchant *model.Merchant) error
	GetAll() ([]model.Merchant, error)
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db),
		Merchant: NewMerchantRepository(db),
		Page:     NewPageRepository(db),
	}
}

type Repositories struct {
	User     UserRepository
	Merchant MerchantRepository
	Page     PageRepository
}
