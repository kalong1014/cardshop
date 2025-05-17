package model

import (
	"time"
)

// 页面模板模型
type PageTemplate struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	MerchantID uint      `gorm:"not null" json:"merchant_id"`  // 所属商户
	Name       string    `gorm:"not null" json:"name"`         // 模板名称
	LayoutData string    `gorm:"type:text" json:"layout_data"` // JSON格式布局数据
	IsDefault  bool      `gorm:"default:false" json:"is_default"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
