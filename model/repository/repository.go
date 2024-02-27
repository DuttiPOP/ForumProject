package repository

import (
	"ForumProject/model/entity"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type IUserRepository interface {
	Create(user entity.User) (int, error)
	Get(id int) (user entity.User, err error)
	Delete(id int) error
}

type Repository struct {
	IUserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewUserRepository(db),
	}
}
