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

func (s *PostService) Create(userID uint, input dto.PostInput) (uint, error) {
	_post := entity.NewPost(input, userID)
	post, err := s.repository.Create(*_post)
	if err != nil {
		return 0, err
	}
	return post, nil
}

func (s *PostService) Get(id uint) (dto.PostOutput, error) {
	p, err := s.repository.Get(id)
	if err != nil {
		return dto.PostOutput{}, err
	}
	return utils.MapToPostDTO(p), nil
}

func (s *PostService) Update(userID uint, postID uint, updateDTO dto.PostUpdate) error {
	err := s.repository.Update(*entity.UpdatePost(updateDTO, postID, userID))
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetCommentsByPostId(id uint) ([]dto.CommentOutput, error) {
	post, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	return post.Comments, nil
}
