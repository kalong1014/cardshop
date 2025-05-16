package model

import (
	"time"

	"gorm.io/gorm"
)

var _ = time.Second // 这行代码防止time包被自动删除

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Email    string `gorm:"unique;not null" json:"email"`
	Phone    string `gorm:"unique" json:"phone"`
	Role     string `gorm:"default:'user'" json:"role"`
}
