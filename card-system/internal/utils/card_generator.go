package utils

import (
    "crypto/rand"
    "encoding/hex"
    "errors"
    "fmt"
    "math/big"
    "strings"
    "time"

    "card-system/internal/common"
)

// CardGenerator 卡密生成器
type CardGenerator struct {
    logger common.Logger
}

// NewCardGenerator 创建卡密生成器
func NewCardGenerator(logger common.Logger) *CardGenerator {
    return &CardGenerator{
        logger: logger,
    }
}

// GenerateCard 生成单个卡密
func (cg *CardGenerator) GenerateCard(length int, prefix string) (string, error) {
    if length <= 0 {
        return "", errors.New("卡密长度必须大于0")
    }
    
    // 计算随机部分长度
    randomLength := length - len(prefix)
    if randomLength <= 0 {
        return "", errors.New("前缀长度不能大于或等于总长度")
    }
    
    // 生成随机字节
    bytes := make([]byte, (randomLength+1)/2)
    _, err := rand.Read(bytes)
    if err != nil {
        cg.logger.Error("生成卡密失败: %v", err)
        return "", err
    }
    
    // 转换为十六进制字符串
    randomPart := hex.EncodeToString(bytes)[:randomLength]
    
    // 添加分隔符
    formattedCard := addSeparators(prefix + randomPart)
    
    cg.logger.Info("生成新卡密: %s", formattedCard)
    return formattedCard, nil
}

// GenerateBatch 批量生成卡密
func (cg *CardGenerator) GenerateBatch(count int, length int, prefix string) ([]string, error) {
    if count <= 0 {
        return nil, errors.New("生成数量必须大于0")
    }
    
    cards := make([]string, 0, count)
    for i := 0; i < count; i++ {
        card, err := cg.GenerateCard(length, prefix)
        if err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }
    
    return cards, nil
}

// addSeparators 添加分隔符，提高可读性
func addSeparators(card string) string {
    // 每4个字符添加一个连字符
    var result strings.Builder
    for i, r := range card {
        if i > 0 && i%4 == 0 {
            result.WriteRune('-')
        }
        result.WriteRune(r)
    }
    return result.String()
}

// ValidateCard 验证卡密格式
func (cg *CardGenerator) ValidateCard(card string) bool {
    // 移除分隔符
    cleanCard := strings.ReplaceAll(card, "-", "")
    
    // 检查长度和字符
    if len(cleanCard) < 8 || len(cleanCard) > 32 {
        return false
    }
    
    // 检查是否为十六进制字符
    for _, r := range cleanCard {
        if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
            return false
        }
    }
    
    return true
}

// GenerateExpiryDate 生成卡密过期日期
func (cg *CardGenerator) GenerateExpiryDate(days int) time.Time {
    return time.Now().AddDate(0, 0, days)
}    