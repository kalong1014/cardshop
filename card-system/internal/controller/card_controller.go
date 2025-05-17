package controller

import (
	"card-system/models"
	"card-system/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 商户生成卡密（需商户权限中间件）
func GenerateCards(c *gin.Context) {
	merchantID, _ := strconv.Atoi(c.Param("merchant_id"))
	var req struct {
		ProductID string `json:"product_id" binding:"required"`
		Count     int    `json:"count" binding:"required,min=1,max=10000"`
		// 可选配置：长度、分段等
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cfg := &utils.CardGenConfig{
		Length:    16,
		Segments:  4,
		Separator: "-",
	}
	cards, err := utils.BatchGenerateCards(uint(merchantID), req.ProductID, req.Count, cfg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "卡密生成失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": len(cards)})
}

// 核销卡密
func RedeemCard(c *gin.Context) {
	var req struct {
		CardCode string `json:"card_code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var card models.Card
	if result := utils.DB.Where("card_code = ?", req.CardCode).First(&card); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "卡密不存在或已过期"})
		return
	}

	if card.Status != models.CardStatusUnused {
		c.JSON(http.StatusBadRequest, gin.H{"error": "卡密已使用或过期"})
		return
	}

	// 标记为已使用
	if err := utils.DB.Model(&card).Updates(map[string]interface{}{
		"status":  models.CardStatusUsed,
		"used_at": time.Now(),
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "核销失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "核销成功"})
}

func GetCards(c *gin.Context) {
	var query models.CardQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	merchantID := getCurrentMerchantID(c)
	query.MerchantID = merchantID
	
	cards, total, err := service.Card.GetCards(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cards":  cards,
		"total":  total,
		"page":   query.Page,
		"pageSize": query.PageSize,
	})
}

// 批量失效卡密
func BatchInvalidateCards(c *gin.Context) {
	var req struct {
		CardIDs []uint `json:"card_ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.BatchUpdateCardStatus(req.CardIDs, models.CardStatusExpired); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "批量操作失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "已批量失效所选卡密"})
}