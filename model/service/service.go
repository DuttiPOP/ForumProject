package service

import (
	"ForumProject/model/repository"
)

type Service struct {
}

func NewService(repository repository.Repository) *Service {
	return &Service{}
}
