package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
	"ForumProject/model/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	emailDuplicate    = "uni_users_email"
	usernameDuplicate = "uni_users_username"
)

var (
	ErrEmailDuplicate    = errors.New("email is already taken")
	ErrUsernameDuplicate = errors.New("username is already taken")
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) Create(input dto.SignUpInput) (dto.UserOutput, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserOutput{}, err
	}
	input.Password = string(hashedPassword)
	user, err := s.repository.Create(*entity.NewUser(input))
	if err != nil {
		if strings.Contains(err.Error(), emailDuplicate) {
			return dto.UserOutput{}, ErrEmailDuplicate
		}
		if strings.Contains(err.Error(), usernameDuplicate) {
			return dto.UserOutput{}, ErrUsernameDuplicate
		}
	}
	return utils.MapToUserDTO(user), err
}

func (s *UserService) Get(id uint) (dto.UserOutput, error) {
	user, err := s.repository.Get(id)
	if err != nil {
		return dto.UserOutput{}, err
	}
	return utils.MapToUserDTO(user), nil
}

func (s *UserService) Delete(id uint) error {
	return s.repository.Delete(id)
}

func (s *UserService) Update(userID uint, input dto.UserUpdate) error {
	user := entity.UpdateUser(userID, input)
	err := s.repository.Update(*user)
	return err
}

func (s *UserService) GetAllPosts(id uint) (posts []dto.PostOutput, err error) {
	user, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	for _, post := range user.Posts {
		posts = append(posts, utils.MapToPostDTO(post))
	}
	return posts, nil
}

func (s *UserService) Authenticate(input dto.SignInInput) (int, error) {
	user, err := s.repository.GetByEmail(input.Email)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return 0, err
	}
	return int(user.ID), nil
}
