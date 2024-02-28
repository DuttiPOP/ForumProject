package service

import (
	"ForumProject/model/entity"
	"ForumProject/model/repository"
)

type IUserService interface {
	Create(user entity.User) (int, error)
	Get(id string) (user entity.User, err error)
	Delete(id string) error
}

type Service struct {
	IUserService
}

func NewService(repository repository.Repository) *Service {
	return &Service{NewUserService(repository.IUserRepository)}
}
