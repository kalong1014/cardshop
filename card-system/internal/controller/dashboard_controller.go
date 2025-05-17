package controller

import (
	"card-system/internal/model"
	"card-system/internal/utils" // 添加此行
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// controller/dashboard_controller.go
func GetDashboardStats(c *gin.Context) {
	var stats struct {
		TotalUsers      int     `json:"total_users"`
		ActiveMerchants int     `json:"active_merchants"`
		TodaySales      float64 `json:"today_sales"`
		TopProduct      string  `json:"top_product"`
	}

	utils.DB.Model(&model.User{}).Count(&stats.TotalUsers)
	utils.DB.Model(&model.Merchant{}).Where("status = ?", model.MerchantStatusApproved).Count(&stats.ActiveMerchants)
	utils.DB.Model(&model.Order{}).Where("created_at >= ?", time.Now().Truncate(24*time.Hour)).Sum("amount", &stats.TodaySales)

	// 查询最畅销产品（按卡密分类）
	var topProduct struct {
		ProductID string `gorm:"product_id"`
		Count     int    `gorm:"count"`
	}
	utils.DB.Model(&model.Order{}).
		Select("product_id, count(*) as count").
		Group("product_id").
		Order("count desc").
		Limit(1).
		Scan(&topProduct)
	stats.TopProduct = topProduct.ProductID

	c.JSON(http.StatusOK, stats)
}
