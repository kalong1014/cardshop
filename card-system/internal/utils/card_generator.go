package utils

import (
	"card-system/internal/model"
	"card-system/internal/utils"
	"math/rand"
	"strings"
	"time"
)

// 卡密生成配置
type CardGenConfig struct {
	Length    int    // 卡密长度
	Segments  int    // 分段数（如XXXX-XXXX格式）
	Separator string // 分隔符
	Chars     string // 可用字符（默认：字母+数字）
}

// 生成单张卡密
func GenerateCardCode(cfg *CardGenConfig) string {
	if cfg.Chars == "" {
		cfg.Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	}
	rand.Seed(time.Now().UnixNano())
	buf := make([]byte, cfg.Length)
	for i := range buf {
		buf[i] = cfg.Chars[rand.Intn(len(cfg.Chars))]
	}
	// 分段处理
	var parts []string
	for i := 0; i < cfg.Segments; i++ {
		start := i * (cfg.Length / cfg.Segments)
		end := start + (cfg.Length / cfg.Segments)
		parts = append(parts, string(buf[start:end]))
	}
	return strings.Join(parts, cfg.Separator)
}

func InitDB() error {
	if err := DB.AutoMigrate(
		&model.User{},
		&model.Merchant{},
		&model.Card{},
		&model.Order{},
	); err != nil {
		return err
	}
	return nil
}

// 批量生成卡密
func BatchGenerateCards(merchantID uint, productID string, count int, cfg *CardGenConfig) ([]*model.Card, error) {
	cards := make([]*model.Card, 0, count)
	for i := 0; i < count; i++ {
		code := GenerateCardCode(cfg)
		card := &model.Card{
			MerchantID: merchantID,
			ProductID:  productID,
			ExpireAt:   time.Now().Add(365 * 24 * time.Hour), // 默认有效期1年
		}
		if err := card.EncryptCode(code); err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, utils.DB.CreateInBatches(cards, 1000).Error // 批量插入
}
