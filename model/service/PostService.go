package service

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
	"ForumProject/model/repository"
	"ForumProject/model/utils"
)

type PostService struct {
	repository repository.IPostRepository
}

func NewPostService(repository repository.IPostRepository) *PostService {
	return &PostService{repository: repository}
}

func (s *PostService) Create(userID uint, input dto.PostInput) (dto.PostOutput, error) {
	post, err := s.repository.Create(*entity.NewPost(userID, input))
	if err != nil {
		return dto.PostOutput{}, err
	}
	return utils.MapToPostDTO(post), nil
}

func (s *PostService) Get(id uint) (dto.PostOutput, error) {
	p, err := s.repository.Get(id)
	if err != nil {
		return dto.PostOutput{}, err
	}
	return utils.MapToPostDTO(p), nil
}

func (s *PostService) Update(userID uint, postID uint, updateDTO dto.PostUpdate) error {
	post, err := s.repository.Get(postID)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return entity.ErrNotOwner
	}
	post.Update(updateDTO)
	return s.repository.Update(post)
}

func (s *PostService) GetCommentsByPostId(id uint) ([]dto.CommentOutput, error) {
	post, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	return post.Comments, nil
}
