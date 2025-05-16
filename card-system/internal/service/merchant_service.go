// internal/service/merchant_service.go
package service

import (
	"card-system/internal/model"
	"card-system/internal/repository"
)

type MerchantService struct {
	merchantRepo repository.MerchantRepository
}

func NewMerchantService(merchantRepo repository.MerchantRepository) *MerchantService {
	return &MerchantService{merchantRepo: merchantRepo}
}

func (s *MerchantService) GetMerchantByID(id uint) (*model.Merchant, error) {
	return s.merchantRepo.FindByID(id)
}

func (s *MerchantService) GetMerchantByUserID(userID uint) (*model.Merchant, error) {
	return s.merchantRepo.FindByUserID(userID)
}

func (s *MerchantService) CreateMerchant(merchant *model.Merchant) error {
	return s.merchantRepo.Create(merchant)
}

func (s *MerchantService) UpdateMerchant(merchant *model.Merchant) error {
	return s.merchantRepo.Update(merchant)
}

func (s *MerchantService) GetAllMerchants() ([]model.Merchant, error) {
	return s.merchantRepo.GetAll()
}
