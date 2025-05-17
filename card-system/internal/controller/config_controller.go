package controller

import (
	"net/http" // 添加此行

	"github.com/gin-gonic/gin"
)

// 获取配置
func GetConfig(c *gin.Context) {
	key := c.Param("key")
	value, err := service.Config.Get(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "配置项不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": value})
}

// 更新配置（需管理员权限）
func UpdateConfig(c *gin.Context) {
	var req struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.Config.Update(req.Key, req.Value); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "配置更新成功"})
}
