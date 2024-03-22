package repository

import (
	"ForumProject/model/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) (entity.User, error)
	Get(id uint) (entity.User, error)
	Delete(id uint) error
	Update(user entity.User) error
	GetByEmail(email string) (entity.User, error)
}

type IPostRepository interface {
	Create(post entity.Post) (entity.Post, error)
	Get(id uint) (entity.Post, error)
	Delete(id uint) error
	Update(post entity.Post) error
}

type ICommentRepository interface {
	Create(comment entity.Comment) (entity.Comment, error)
	Get(id uint) (comment entity.Comment, err error)
	Delete(id uint) error
	Update(id uint, comment entity.Comment) error
}

type Repository struct {
	IUserRepository
	IPostRepository
	ICommentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		NewUserRepository(db),
		NewPostRepository(db),
		NewCommentRepository(db),
	}
}
