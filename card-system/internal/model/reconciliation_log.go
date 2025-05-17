package model

import (
	"gorm.io/gorm"
)

type ReconciliationStatus string

const (
	ReconciliationStatusMatch    ReconciliationStatus = "match"
	ReconciliationStatusMismatch ReconciliationStatus = "mismatch"
)

type ReconciliationLog struct {
	ID         uint                  `gorm:"primaryKey" json:"id"`
	Date       string                `gorm:"not null" json:"date"`
	OrderTotal float64               `gorm:"not null" json:"order_total"`
	BillTotal  float64               `gorm:"not null" json:"bill_total"`
	Difference float64               `json:"difference"`
	Status     ReconciliationStatus  `gorm:"not null" json:"status"`
	CreatedAt  time.Time             `json:"created_at"`
}