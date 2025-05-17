package service

import (
	"card-system/internal/model"
	"card-system/internal/repository"
)

type CardService struct {
	cardRepo repository.CardRepository
}

func NewCardService(cardRepo repository.CardRepository) *CardService {
	return &CardService{cardRepo: cardRepo}
}

func (s *CardService) GenerateCards(merchantID uint, productID uint, count int, length int) ([]model.Card, error) {
	var cards []model.Card
	for i := 0; i < count; i++ {
		cardCode, err := model.GenerateCardCode(length)
		if err != nil {
			return nil, err
		}
		card := model.Card{
			MerchantID: merchantID,
			ProductID:  productID,
			CardCode:   cardCode,
		}
		if err := s.cardRepo.Create(&card); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}
