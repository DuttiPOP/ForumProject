package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
	"errors"
	"strings"
)

const (
	minUsernameLen    = 3
	maxUsernameLen    = 20
	minPasswordLen    = 6
	maxPasswordLen    = 100
	emailDuplicate    = "uni_users_email"
	usernameDuplicate = "uni_users_username"
)

var (
	ErrEmailDuplicate    = errors.New("email is already taken")
	ErrUsernameDuplicate = errors.New("username is already taken")
	ErrInvalidUserId     = errors.New("invalid user id")
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) Create(input dto.SignUpInput) (int, error) {
	user, err := entity.NewUser(input)
	if err != nil {
		return 0, err
	}
	id, err := service.repository.Create(*user)
	if err != nil {
		if strings.Contains(err.Error(), emailDuplicate) {
			return 0, ErrEmailDuplicate
		}
		if strings.Contains(err.Error(), usernameDuplicate) {
			return 0, ErrUsernameDuplicate
		}
	}
	return id, err
}

func (service *UserService) Get(id uint) (entity.User, error) {
	user, err := service.repository.Get(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (service *UserService) Delete(id uint) error {
	return service.repository.Delete(id)
}

func (service *UserService) Update(id uint, input dto.UserUpdateDTO) error {
	u, err := service.repository.Get(id)
	if err != nil {
		return err
	}
	err = u.UpdateUser(input)
	if err != nil {
		return err
	}

	return service.repository.Update(id, u)
}
