package controller

import (
	"card-system/internal/common"
	"card-system/internal/model"
	"card-system/internal/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		common.Error(c, 400, err.Error()) // 使用公共响应函数
		return
	}

	if err := service.UserRegister(&user); err != nil {
		common.Error(c, 500, "注册失败")
		return
	}

	common.Success(c, gin.H{"message": "注册成功"})
}
