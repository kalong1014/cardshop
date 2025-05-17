package controller 


import (
	"card-system/internal/model"
	"card-system/internal/utils"
	"net/http"
    "card-system/internal/utils" // 添加此行
    "card-system/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req struct {
		CardID        uint    `json:"card_id" binding:"required"`
		PaymentMethod string  `json:"payment_method" binding:"required,oneof=alipay wechat test"`
		Amount        float64 `json:"amount" binding:"required,min=0.01"`
		recordOrder(order.Status) // 记录订单状态指标
		c.JSON(http.StatusOK, order)
}
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证卡密是否属于当前商户（需从JWT获取商户ID，此处简化为参数）
	merchantID, _ := strconv.Atoi(c.Param("merchant_id"))

	order := &model.Order{
		MerchantID:   uint(merchantID),
		UserID:       getCurrentUserID(c), // 从JWT获取用户ID
		CardID:       req.CardID,
		Amount:       req.Amount,
		PaymentMethod: req.PaymentMethod,
	}
	if err := utils.DB.Create(order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "订单创建失败"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func getCurrentUserID(c *gin.Context) uint {
	claims := c.MustGet("user").(jwt.MapClaims)
	return uint(claims["id"].(float64))
}

// 取消订单
func CancelOrder(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))
	var order model.Order
	if err := utils.DB.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "订单不存在"})
		return
	}
	if err := order.Cancel(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "订单已取消"})
}

// 获取订单列表
func GetOrders(c *gin.Context) {
	var query model.OrderQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 商户只能查看自己的订单（普通用户查看自己的订单，商户查看旗下订单）
	if isMerchant := c.GetBool("is_merchant"); isMerchant {
		query.MerchantID = getCurrentMerchantID(c)
	} else {
		query.UserID = getCurrentUserID(c)
	}
	
	orders, total, err := service.Order.GetOrders(&query)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders, "total": total})
}
