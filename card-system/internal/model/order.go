package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	MerchantID    uint       `gorm:"not null" json:"merchant_id"`
	UserID        uint       `gorm:"not null" json:"user_id"`
	CardID        uint       `json:"card_id"`
	Amount        float64    `gorm:"not null" json:"amount"`
	PaymentMethod string     `json:"payment_method"`
	TransactionID string     `json:"transaction_id"`
	Status        string     `gorm:"default:'pending'" json:"status"`
	PaidAt        *time.Time `json:"paid_at"`
	Merchant      Merchant   `gorm:"foreignKey:MerchantID" json:"merchant"`
	User          User       `gorm:"foreignKey:UserID" json:"user"`
	Card          Card       `gorm:"foreignKey:CardID" json:"card"`
}
