package repository

import (
	"card-system/internal/model"
	"card-system/internal/utils" // 仅依赖工具包的数据库连接
)

func CreateUser(user *model.User) error {
	return utils.DB.Create(user).Error // 直接使用 utils 中的数据库连接
}
