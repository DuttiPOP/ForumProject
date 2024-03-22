package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/repository"
)

type IUserService interface {
	Create(user dto.SignUpInput) (dto.UserOutput, error)
	Get(id uint) (dto.UserOutput, error)
	Delete(id uint) error
	Update(id uint, input dto.UserUpdate) error
	GetAllPosts(id uint) ([]dto.PostOutput, error)
	Authenticate(input dto.SignInInput) (int, error)
}

type IPostService interface {
	Create(userID uint, input dto.PostInput) (dto.PostOutput, error)
	Get(id uint) (dto.PostOutput, error)
	Update(userID uint, postID uint, updateDTO dto.PostUpdate) error
	GetCommentsByPostId(id uint) ([]dto.CommentOutput, error)
}

type ICommentService interface {
	Create(userID uint, postID uint, input dto.CommentInput) (dto.CommentOutput, error)
	Update(userID uint, commentID uint, updateDTO dto.CommentUpdate) error
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
