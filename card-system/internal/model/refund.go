package model

import (
	"time"
)

// 退款状态枚举
type RefundStatus string

const (
	RefundStatusPending  RefundStatus = "pending"  // 待审核
	RefundStatusApproved RefundStatus = "approved" // 已同意
	RefundStatusRejected RefundStatus = "rejected" // 已拒绝
)

// 退款申请
type Refund struct {
	ID         uint         `gorm:"primaryKey" json:"id"`
	OrderID    uint         `gorm:"not null" json:"order_id"`
	UserID     uint         `gorm:"not null" json:"user_id"`
	MerchantID uint         `gorm:"not null" json:"merchant_id"`
	Amount     float64      `gorm:"not null" json:"amount"`
	Reason     string       `gorm:"type:text" json:"reason"`
	Status     RefundStatus `gorm:"default:pending" json:"status"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}
