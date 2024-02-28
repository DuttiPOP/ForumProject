package repository

import (
	"ForumProject/model/entity"
	"gorm.io/gorm"
)

const (
	userTable = "users"
)

type IUserRepository interface {
	Create(user entity.User) (int, error)
	Get(id uint) (user entity.User, err error)
	Delete(id uint) error
	Update(id uint, user entity.User) error
}

type Repository struct {
	IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewUserRepository(db),
	}
}
