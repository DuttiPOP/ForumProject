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

func (r *UserRepository) Create(user entity.User) (entity.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
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

func (r *UserRepository) Update(user entity.User) error {
	result := r.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(user)
	return result.Error
}

func (r *UserRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
