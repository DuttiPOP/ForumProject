package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
	"ForumProject/model/utils"
)

type CommentService struct {
	repository repository.ICommentRepository
}

func NewCommentService(repository repository.ICommentRepository) *CommentService {
	return &CommentService{repository: repository}
}

func (s *CommentService) Create(userID uint, postID uint, input dto.CommentInput) (dto.CommentOutput, error) {
	comment, err := s.repository.Create(*entity.NewComment(userID, postID, input))
	if err != nil {
		return dto.CommentOutput{}, err
	}
	return utils.MapToCommentDTO(comment), nil
}

func (s *CommentService) Update(userID uint, commentID uint, updateDTO dto.CommentUpdate) error {
	comment, err := s.repository.Get(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID {
		return entity.ErrNotOwner
	}
	comment.Update(updateDTO)
	return s.repository.Update(commentID, comment)
}
