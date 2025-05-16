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

func (s *PageService) CreatePage(merchantID uint, name string) (*model.Page, error) {
	page := &model.Page{
		MerchantID: merchantID,
		Name:       name,
	}
	err := s.pageRepo.Create(page)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *PageService) UpdatePage(page *model.Page) error {
	return s.pageRepo.Update(page)
}

func (s *PageService) DeletePage(id uint) error {
	return s.pageRepo.Delete(id)
}

func (s *PageService) GetMerchantPages(merchantID uint) ([]model.Page, error) {
	return s.pageRepo.GetByMerchantID(merchantID)
}
