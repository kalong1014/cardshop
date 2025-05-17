package card

import (
	"crypto/rand"
	"math/big"
	"strings"
	"time"
)

type CardGenerator struct {
	// 可以添加配置信息
}

func NewCardGenerator() *CardGenerator {
	return &CardGenerator{}
}

func (cg *CardGenerator) GenerateCardNumber(cardType string) (string, error) {
	var prefix string

	switch cardType {
	case "visa":
		prefix = "4"
	case "mastercard":
		prefix = "5"
	case "amex":
		prefix = "37"
	default:
		prefix = "6" // 自定义卡类型
	}

	// 生成随机数部分
	remainingDigits := 16 - len(prefix)
	randomPart := make([]byte, remainingDigits)

	for i := 0; i < remainingDigits; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		randomPart[i] = byte(num.Int64() + 48) // 48是ASCII码的'0'
	}

	cardNumber := prefix + string(randomPart)

	// 应用Luhn算法验证
	if !cg.validateLuhn(cardNumber) {
		// 如果验证失败，尝试修复或重新生成
		return cg.fixLuhn(cardNumber)
	}

	// 格式化卡号，每4位分隔
	return cg.formatCardNumber(cardNumber), nil
}

func (cg *CardGenerator) validateLuhn(cardNumber string) bool {
	// 实现Luhn算法验证
	// 这里简化处理，实际应用中应该完整实现
	return true
}

func (cg *CardGenerator) fixLuhn(cardNumber string) (string, error) {
	// 修复Luhn算法校验失败的卡号
	// 简化处理
	return cardNumber, nil
}

func (cg *CardGenerator) formatCardNumber(cardNumber string) string {
	var result strings.Builder

	for i, char := range cardNumber {
		if i > 0 && i%4 == 0 {
			result.WriteByte(' ')
		}
		result.WriteRune(char)
	}

	return result.String()
}

func (cg *CardGenerator) GenerateCVV() (string, error) {
	// 生成3位CVV码
	cvv := make([]byte, 3)

	for i := 0; i < 3; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		cvv[i] = byte(num.Int64() + 48)
	}

	return string(cvv), nil
}

func (cg *CardGenerator) GenerateExpiryDate() string {
	// 生成有效期（格式：MM/YY）
	now := time.Now()
	expiry := now.AddDate(3, 0, 0)
	return expiry.Format("01/06")
}
