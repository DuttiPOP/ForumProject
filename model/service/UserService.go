package service

import (
	"ForumProject/model/entity"
	"ForumProject/model/repository"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) Create(user entity.User) (int, error) {
	return service.repository.Create(user)
}

func (service *UserService) Get(id int) (user entity.User, err error) {
	return service.repository.Get(id)
}

func (service *UserService) Delete(id int) error {
	return service.repository.Delete(id)
}
