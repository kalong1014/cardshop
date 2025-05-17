package middleware

import (
	"card-system/internal/model"
	"card-system/internal/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		// 记录登录日志
		if c.Request.URL.Path == "/api/login" {
			var log model.OperationLog
			log.UserID = getCurrentUserID(c)
			log.Username = c.GetString("username")
			log.Type = model.LogTypeLogin
			log.Action = fmt.Sprintf("登录成功，状态码：%d", c.Writer.Status())
			log.IP = getClientIP(c)
			utils.DB.Create(&log)
		}

		// 记录其他操作日志（示例：商户操作）
		if strings.HasPrefix(c.Request.URL.Path, "/api/merchant/") {
			// 解析商户ID和操作类型
			// ... 具体逻辑根据路由设计补充 ...
		}
	}
}

func getClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
