package model

import (
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	MerchantID uint   `gorm:"not null" json:"merchant_id"`
	Name       string `gorm:"not null" json:"name"`
	Status     string `gorm:"default:'draft'" json:"status"`
	Elements   string `gorm:"type:text" json:"elements"` // 页面元素JSON
	User       User   `gorm:"foreignKey:MerchantID" json:"user"`
}
