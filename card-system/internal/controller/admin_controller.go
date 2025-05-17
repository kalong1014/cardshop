// 控制器（controllers/admin_controller.go）
func GetSystemStats(c *gin.Context) {
	var stats struct {
		MerchantCount int `json:"merchant_count"`
		CardCount     int `json:"card_count"`
		TodayOrders   int `json:"today_orders"`
	}
	utils.DB.Model(&models.Merchant{}).Count(&stats.MerchantCount)
	utils.DB.Model(&models.Card{}).Count(&stats.CardCount)
	utils.DB.Model(&models.Order{}).Where("created_at >= ?", 
		time.Now().Truncate(24*time.Hour)).Count(&stats.TodayOrders)
	c.JSON(http.StatusOK, stats)
}