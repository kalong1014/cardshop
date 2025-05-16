// internal/service/page_service.go
package service

import (
	"card-system/internal/model"
	"card-system/internal/repository"
)

type PageService struct {
	pageRepo repository.PageRepository
}

func NewPageService(pageRepo repository.PageRepository) *PageService {
	return &PageService{pageRepo: pageRepo}
}

func (s *PageService) GetPageByID(id uint) (*model.Page, error) {
	return s.pageRepo.GetByID(id)
}

func (s *PageService) CreatePage(page *model.Page) error {
	return s.pageRepo.Create(page)
}

func (s *PageService) UpdatePage(page *model.Page) error {
	return s.pageRepo.Update(page)
}

func (s *PageService) DeletePage(id uint) error {
	return s.pageRepo.Delete(id)
}

func (s *PageService) GetPagesByMerchantID(merchantID uint) ([]model.Page, error) {
	return s.pageRepo.GetByMerchantID(merchantID)
}
