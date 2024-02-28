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

func (repository *UserRepository) Create(user entity.User) (int, error) {
	result := repository.db.Table(userTable).Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

func (repository *UserRepository) Get(id uint) (entity.User, error) {
	var user entity.User
	result := repository.db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (repository *UserRepository) Delete(id uint) error {
	result := repository.db.Delete(&entity.User{}, id)
	return result.Error
}

func (repository *UserRepository) Update(id uint, user entity.User) error {
	return repository.db.Model(&user).Where("id = ?", id).Updates(&user).Error
}
