package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 完整错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":  "服务器内部错误",
					"detail": fmt.Sprintf("%v", err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
