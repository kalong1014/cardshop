package service

import (
	"card-system/internal/model"

	"gorm.io/gorm"
)

func (s *ConfigService) Get(key string) (string, error) {
	var config model.SystemConfig
	if err := utils.DB.Where("key = ?", key).First(&config).Error; err != nil {
		return "", err
	}
	return config.Value, nil
}

func (s *ConfigService) Update(key, value string) error {
	return utils.DB.Transaction(func(tx *gorm.DB) error {
		var config model.SystemConfig
		if err := tx.Where("key = ?", key).First(&config).Error; err != nil {
			return tx.Create(&model.SystemConfig{Key: key, Value: value}).Error
		}
		return tx.Model(&config).Update("value", value).Error
	})
}
