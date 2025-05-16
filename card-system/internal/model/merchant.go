package model

import (
	"time"

	"gorm.io/gorm"
)

var _ = time.Second // 这行代码防止time包被自动删除

type Merchant struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
	Logo   string `json:"logo"`
	Status string `gorm:"default:'pending'" json:"status"`
	Level  string `gorm:"default:'basic'" json:"level"`
	Domain string `gorm:"unique" json:"domain"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
