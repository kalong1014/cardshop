// internal/service/user_service.go
package service

import (
	"card-system/internal/model"
	"card-system/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.userRepo.Update(user)
}
