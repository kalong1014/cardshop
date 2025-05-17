package controller

import (
	"card-system/internal/model"
	"card-system/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// controller/stats_controller.go
func GetDailyOrderTrend(c *gin.Context) {
	// 获取近30天数据
	dates := make([]string, 30)
	orders := make([]float64, 30)
	for i := 0; i < 30; i++ {
		date := time.Now().Add(-24 * time.Hour * time.Duration(i)).Format("2006-01-02")
		dates[29-i] = date // 倒序排列
		var count int64
		utils.DB.Model(&model.Order{}).Where("created_at >= ? AND created_at < ?", date, date+" 23:59:59").Count(&count)
		orders[29-i] = float64(count)
	}
	c.JSON(http.StatusOK, gin.H{"dates": dates, "orders": orders})
}
