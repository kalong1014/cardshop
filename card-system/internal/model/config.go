package model

// 系统配置
type SystemConfig struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Key    string `gorm:"uniqueIndex" json:"key"` // 配置键（如payment_channel.alipay.app_id）
	Value  string `gorm:"type:text" json:"value"` // 配置值
	Remark string `json:"remark"`
}
