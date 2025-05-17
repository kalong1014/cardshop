package service

import (
	"errors"
	"sync"
	"time"

	"card-system/internal/card"
	"card-system/internal/cert"
	"card-system/internal/model"
	"card-system/internal/repository"
)

type CardService struct {
	repo     repository.CardRepository
	redis    *redis.Client
	logger   *common.DefaultLogger // 确保使用 common 日志
}

func NewCardService(repo repository.CardRepository, redis *redis.Client, logger *common.DefaultLogger) *CardService {
	return &CardService{
		repo:     repo,
		redis:    redis,
		logger:   logger,
	}
}

func (cs *CardService) GenerateCard(userID int, cardType string) (*model.Card, error) {
	cs.cardLocker.Lock()
	defer cs.cardLocker.Unlock()

	user, err := cs.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if !user.IsActive {
		return nil, errors.New("用户未激活")
	}

	// 检查用户卡片数量限制
	cardCount, err := cs.cardRepo.GetCardCountByUser(userID)
	if err != nil {
		return nil, errors.New("获取用户卡片数量失败")
	}

	if cardCount >= 5 {
		return nil, errors.New("卡片数量已达上限")
	}

	// 生成卡片编号
	cardNumber, err := cs.cardGen.GenerateCardNumber(cardType)
	if err != nil {
		return nil, errors.New("生成卡片编号失败")
	}

	// 生成安全码
	cvv, err := cs.cardGen.GenerateCVV()
	if err != nil {
		return nil, errors.New("生成安全码失败")
	}

	// 生成有效期
	expiryDate := time.Now().AddDate(3, 0, 0)

	newCard := &model.Card{
		UserID:     userID,
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpiryDate: expiryDate,
		CardType:   cardType,
		Status:     "active",
	}

	// 保存卡片到数据库
	err = cs.cardRepo.SaveCard(newCard)
	if err != nil {
		return nil, errors.New("保存卡片失败")
	}

	// 生成卡片证书
	certData, err := cs.certMgr.GenerateCardCert(newCard.CardNumber)
	if err != nil {
		// 回滚卡片创建
		_ = cs.cardRepo.DeleteCard(newCard.ID)
		return nil, errors.New("生成卡片证书失败")
	}

	newCard.CertData = certData

	return newCard, nil
}

func (cs *CardService) GetCardByNumber(cardNumber string) (*model.Card, error) {
	card, err := cs.cardRepo.GetCardByNumber(cardNumber)
	if err != nil {
		return nil, errors.New("卡片不存在")
	}

	// 验证证书
	isValid, err := cs.certMgr.VerifyCardCert(card.CardNumber, card.CertData)
	if err != nil || !isValid {
		return nil, errors.New("卡片证书验证失败")
	}

	return card, nil
}

func (cs *CardService) FreezeCard(cardID int) error {
	card, err := cs.cardRepo.GetCardByID(cardID)
	if err != nil {
		return errors.New("卡片不存在")
	}

	if card.Status == "frozen" {
		return nil
	}

	card.Status = "frozen"
	return cs.cardRepo.UpdateCard(card)
}
