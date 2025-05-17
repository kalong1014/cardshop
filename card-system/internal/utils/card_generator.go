package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"card-system/internal/model"
	"card-system/pkg/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CardGenerator 卡片生成器接口
type CardGenerator interface {
	GenerateCardNumber(prefix string, length int) (string, error)
	GenerateCardExpiry(years int) string
	GenerateCVV() (string, error)
	CreatePhysicalCard(userID, merchantID uint, cardType string) (*model.Card, error)
	CreateVirtualCard(userID, merchantID uint, cardType string) (*model.Card, error)
}

// CardGeneratorImpl 卡片生成器实现
type CardGeneratorImpl struct {
	db *gorm.DB
}

// NewCardGenerator 创建新的卡片生成器
func NewCardGenerator(db *gorm.DB) CardGenerator {
	return &CardGeneratorImpl{db: db}
}

// GenerateCardNumber 生成卡号，符合Luhn算法校验
func (cg *CardGeneratorImpl) GenerateCardNumber(prefix string, length int) (string, error) {
	if len(prefix) >= length {
		return "", fmt.Errorf("prefix length cannot be greater than or equal to total length")
	}

	// 生成随机数部分
	remainingDigits := length - len(prefix) - 1 // 减去前缀长度和校验位
	randomPart := make([]byte, remainingDigits)

	for i := 0; i < remainingDigits; i++ {
		// 生成0-9的随机数字
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		randomPart[i] = byte(n.Int64()) + '0'
	}

	// 组合前缀和随机部分
	cardNumber := prefix + string(randomPart)

	// 计算校验位
	checkDigit := calculateLuhnCheckDigit(cardNumber)
	cardNumber += string(checkDigit)

	// 验证卡号是否已存在
	var count int64
	err := cg.db.Model(&model.Card{}).Where("card_number = ?", cardNumber).Count(&count).Error
	if err != nil {
		return "", err
	}

	if count > 0 {
		// 如果已存在，递归重新生成
		return cg.GenerateCardNumber(prefix, length)
	}

	return cardNumber, nil
}

// calculateLuhnCheckDigit 计算Luhn算法校验位
func calculateLuhnCheckDigit(number string) byte {
	sum := 0
	parity := len(number) % 2

	for i, c := range number {
		digit := int(c - '0')

		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	checkDigit := (10 - (sum % 10)) % 10
	return byte(checkDigit) + '0'
}

// GenerateCardExpiry 生成卡片有效期 (MM/YY)
func (cg *CardGeneratorImpl) GenerateCardExpiry(years int) string {
	expiryDate := time.Now().AddDate(years, 0, 0)
	return fmt.Sprintf("%02d/%02d", expiryDate.Month(), expiryDate.Year()%100)
}

// GenerateCVV 生成3位CVV码
func (cg *CardGeneratorImpl) GenerateCVV() (string, error) {
	cvvBytes := make([]byte, 2)
	_, err := rand.Read(cvvBytes)
	if err != nil {
		return "", err
	}

	// 将字节转换为3位数字
	cvvInt := int(cvvBytes[0])<<8 | int(cvvBytes[1])
	cvv := fmt.Sprintf("%03d", cvvInt%1000)

	return cvv, nil
}

// CreatePhysicalCard 创建实体卡
func (cg *CardGeneratorImpl) CreatePhysicalCard(userID, merchantID uint, cardType string) (*model.Card, error) {
	// 开始事务
	return cg.createCard(userID, merchantID, cardType, true)
}

// CreateVirtualCard 创建虚拟卡
func (cg *CardGeneratorImpl) CreateVirtualCard(userID, merchantID uint, cardType string) (*model.Card, error) {
	// 开始事务
	return cg.createCard(userID, merchantID, cardType, false)
}

// createCard 创建卡片的通用逻辑
func (cg *CardGeneratorImpl) createCard(userID, merchantID uint, cardType string, isPhysical bool) (*model.Card, error) {
	var prefix string
	var length int

	// 根据卡类型设置前缀和长度
	switch strings.ToLower(cardType) {
	case "visa":
		prefix = "4"
		length = 16
	case "mastercard":
		prefix = "51"
		length = 16
	case "amex":
		prefix = "34"
		length = 15
	default:
		return nil, fmt.Errorf("unsupported card type: %s", cardType)
	}

	// 生成卡号
	cardNumber, err := cg.GenerateCardNumber(prefix, length)
	if err != nil {
		return nil, fmt.Errorf("failed to generate card number: %v", err)
	}

	// 生成有效期 (5年后)
	expiry := cg.GenerateCardExpiry(5)

	// 生成CVV
	cvv, err := cg.GenerateCVV()
	if err != nil {
		return nil, fmt.Errorf("failed to generate CVV: %v", err)
	}

	// 生成卡片密钥
	cardKey := generateCardKey()

	// 创建卡片实例
	card := &model.Card{
		UserID:       userID,
		MerchantID:   merchantID,
		CardType:     cardType,
		CardNumber:   cardNumber,
		ExpiryDate:   expiry,
		CVV:          cvv,
		CardKey:      cardKey,
		IsPhysical:   isPhysical,
		IsActive:     true,
		ActivationAt: time.Now(),
		UUID:         uuid.New().String(),
	}

	// 在事务中保存卡片
	err = cg.db.Transaction(func(tx *gorm.DB) error {
		// 保存卡片
		if err := tx.Create(card).Error; err != nil {
			logger.Errorf("Failed to create card: %v", err)
			return err
		}

		// 如果是实体卡，记录制卡日志
		if isPhysical {
			cardLog := &model.CardLog{
				CardID:    card.ID,
				Event:     "card_creation",
				Details:   fmt.Sprintf("Physical card created: %s", cardNumber),
				IPAddress: "system",
			}

			if err := tx.Create(cardLog).Error; err != nil {
				logger.Errorf("Failed to create card log: %v", err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return card, nil
}

// generateCardKey 生成卡片密钥
func generateCardKey() string {
	keyBytes := make([]byte, 16)
	rand.Read(keyBytes)
	return hex.EncodeToString(keyBytes)
}
