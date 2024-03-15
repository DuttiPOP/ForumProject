package repository

import (
	"ForumProject/model/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) (int, error)
	Get(id uint) (user entity.User, err error)
	Delete(id uint) error
	Update(id uint, user entity.User) error
	GetByEmail(email string) (user entity.User, err error)
}

type IPostRepository interface {
	Create(post entity.Post) (uint, error)
	Get(id uint) (post entity.Post, err error)
	Delete(id uint) error
	Update(post entity.Post) error
}

type ICommentRepository interface {
	Create(comment entity.Comment) (uint, error)
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
