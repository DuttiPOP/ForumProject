package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
)

type CommentService struct {
	repository repository.ICommentRepository
}

func NewCommentService(repository repository.ICommentRepository) *CommentService {
	return &CommentService{repository: repository}
}

func (s *CommentService) Create(userID uint, postID uint, input dto.CommentInput) (uint, error) {
	newComment := entity.NewComment(input)
	newComment.PostID = postID
	newComment.UserID = userID
	id, err := s.repository.Create(*newComment)
	if err != nil {
		return 0, err
	}
	return id, nil
}
