package model

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	MerchantID uint       `gorm:"not null" json:"merchant_id"`
	ProductID  uint       `json:"product_id"`
	CardCode   string     `gorm:"not null" json:"card_code"`
	Status     string     `gorm:"default:'unused'" json:"status"`
	ExpireAt   *time.Time `json:"expire_at"`
	UsedAt     *time.Time `json:"used_at"`
	Merchant   Merchant   `gorm:"foreignKey:MerchantID" json:"merchant"`
}
