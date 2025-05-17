// 控制器（controller/admin_controller.go）
package controller

import (
	"card-system/internal/model"
	"card-system/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSystemStats(c *gin.Context) {
	var stats struct {
		MerchantCount int `json:"merchant_count"`
		CardCount     int `json:"card_count"`
		TodayOrders   int `json:"today_orders"`
	}
	utils.DB.Model(&model.Merchant{}).Count(&stats.MerchantCount)
	utils.DB.Model(&model.Card{}).Count(&stats.CardCount)
	utils.DB.Model(&model.Order{}).Where("created_at >= ?",
		time.Now().Truncate(24*time.Hour)).Count(&stats.TodayOrders)
	c.JSON(http.StatusOK, stats)
}
