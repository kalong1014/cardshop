package controller

import (
	"card-system/model"
	"card-system/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 支付回调接口（幂等性处理）
func PaymentCallback(c *gin.Context) {
	transactionID := c.Query("transaction_id")
	// 检查回调是否已处理（使用Redis缓存）
	key := fmt.Sprintf("payment:callback:%s", transactionID)
	if utils.Redis.Exists(context.Background(), key).Val() == 1 {
		c.JSON(http.StatusOK, "success") // 重复回调直接返回成功
		return
	}

	var order model.Order
	if result := utils.DB.Where("transaction_id = ?", transactionID).First(&order); result.Error != nil {
		c.JSON(http.StatusNotFound, "订单不存在")
		return
	}

	// 处理订单状态更新
	if err := utils.DB.Transaction(func(tx *gorm.DB) error {
		if order.Status != model.OrderStatusPending {
			return nil // 状态已变更，无需处理
		}
		return tx.Model(&order).Updates(map[string]interface{}{
			"status":    models.OrderStatusPaid,
			"paid_at":   time.Now(),
			"transaction_id": transactionID,
		}).Error
	}); err != nil {
		c.JSON(http.StatusInternalServerError, "处理失败")
		return
	}

	// 记录回调已处理
	utils.Redis.Set(context.Background(), key, "processed", 24*time.Hour)
	c.JSON(http.StatusOK, "success")
}