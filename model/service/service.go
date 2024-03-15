package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
)

type IUserService interface {
	Create(user dto.SignUpInput) (int, error)
	Get(id uint) (user entity.User, err error)
	Delete(id uint) error
	Update(id uint, input dto.UserUpdate) error
	GetAllPosts(id uint) ([]dto.PostOutput, error)
	Authenticate(input dto.SignInInput) (int, error)
}

type IPostService interface {
	Create(userID uint, input dto.PostInput) (uint, error)
	Get(id uint) (dto.PostOutput, error)
	Update(userID uint, postID uint, updateDTO dto.PostUpdate) error
	GetCommentsByPostId(id uint) ([]dto.CommentOutput, error)
}

type ICommentService interface {
	Create(userID uint, postID uint, input dto.CommentInput) (uint, error)
}

type Service struct {
	IUserService
	IPostService
	ICommentService
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		NewUserService(repository.IUserRepository),
		NewPostService(repository.IPostRepository),
		NewCommentService(repository.ICommentRepository),
	}
}
