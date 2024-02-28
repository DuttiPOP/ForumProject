package service

import (
	"ForumProject/model/entity"
	"ForumProject/model/repository"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	minUsernameLen = 3
	maxUsernameLen = 20
	minPasswordLen = 6
	maxPasswordLen = 100
)

var (
	ErrEmailRequired     = errors.New("field \"email\" is required")
	ErrInvalidEmail      = errors.New("field \"email\" is invalid")
	ErrUsernameRequired  = errors.New(fmt.Sprintf("field \"username\" is required and must be between %d and %d characters", minUsernameLen, maxUsernameLen))
	ErrPasswordRequired  = errors.New(fmt.Sprintf("field \"password\" is required and must be between %d and %d characters", minPasswordLen, maxPasswordLen))
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

func (service *UserService) Create(user entity.User) (int, error) {
	return service.repository.Create(user)
}

func ValidateUserFields(user *entity.User) error {
	if user.Email == "" {
		return ErrEmailRequired
	}
	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}
	if len(user.Username) < minUsernameLen || len(user.Username) > maxUsernameLen {
		return ErrUsernameRequired
	}
	if len(user.Password) < minPasswordLen || len(user.Password) > maxPasswordLen {
		return ErrPasswordRequired
	}
	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}
