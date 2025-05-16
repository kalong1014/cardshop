package repository

import "gorm.io/gorm"

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db),
		Merchant: NewMerchantRepository(db),
		Page:     NewPageRepository(db),
	}
}

type Repositories struct {
	User     *UserRepository
	Merchant *MerchantRepository
	Page     *PageRepository
}
