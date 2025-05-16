package repository

import (
	"card-system/internal/model"
)

func (r *UserRepositoryImpl) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) Update(user *model.User) error {
	return r.db.Save(user).Error
}
