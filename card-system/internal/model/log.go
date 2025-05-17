package model

import (
	"gorm.io/gorm"
	"time"
)

// 操作日志类型
type LogType string

const (
	LogTypeLogin      LogType = "login"     // 登录
	LogTypeMerchant   LogType = "merchant"  // 商户管理
	LogTypeCard       LogType = "card"      // 卡密操作
	LogTypePayment    LogType = "payment"   // 支付相关
)

// 操作日志
type OperationLog struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	Username  string     `json:"username"`
	Type      LogType    `gorm:"not null" json:"type"`
	Action    string     `gorm:"not null" json:"action"`
	IP        string     `json:"ip"`
	CreatedAt time.Time  `json:"created_at"`
}