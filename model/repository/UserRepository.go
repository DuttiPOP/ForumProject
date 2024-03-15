package repository

import (
	"ForumProject/model/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user entity.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

func (r *UserRepository) Get(id uint) (entity.User, error) {
	var user entity.User
	result := r.db.Preload("Posts").First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.User{}, id)
	return result.Error
}

func (r *UserRepository) Update(id uint, user entity.User) error {
	return r.db.Model(&user).Where("id = ?", id).Updates(&user).Error
}

func (r *UserRepository) GetByEmail(email string) (user entity.User, err error) {
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
