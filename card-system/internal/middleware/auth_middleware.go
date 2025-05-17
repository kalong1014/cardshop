package middleware

import (
	"card-system/internal/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			common.Error(c, 401, "缺少认证令牌") // 使用公共响应函数
			c.Abort()
			return
		}

		// 解析 JWT（示例逻辑，需结合实际认证逻辑）
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			common.Error(c, 401, "无效的认证令牌")
			c.Abort()
			return
		}

		c.Next()
	}
}
