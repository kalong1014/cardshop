package model

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"gorm.io/gorm"
)

// 卡片状态
const (
	CardStatusUnused   = "unused"
	CardStatusUsed     = "used"
	CardStatusExpired  = "expired"
	CardStatusRefunded = "refunded"
)

// 订单状态
const (
	OrderStatusPending  = "pending"
	OrderStatusPaid     = "paid"
	OrderStatusRefunded = "refunded"
)

// 商户状态
const (
	MerchantStatusApproved = "approved"
	MerchantStatusPending  = "pending"
)

type Card struct {
	gorm.Model
	CardNumber  string `gorm:"uniqueIndex"` // 添加字段
	CVV         string
	ExpiryDate  time.Time
	CardType    string // 类型（如虚拟卡、实体卡）
	Status      string // 状态（有效/已使用/过期）
	MerchantID  uint   // 所属商户
	UserID      uint   // 购买用户
}

// GenerateCardCode 生成卡密
func GenerateCardCode(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
