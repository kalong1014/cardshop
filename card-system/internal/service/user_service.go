package service

import (
	"card-system/internal/model"
	"card-system/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(user *model.User) error {
	// 密码加密（原 utils 逻辑移至此）
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	// 调用数据访问层
	return repository.CreateUser(user)
}
