package repository

import (
	"ForumProject/model/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post entity.Post) (uint, error) {
	result := r.db.Create(&post)
	if result.Error != nil {
		return 0, result.Error
	}

	return post.ID, nil
}

func (r *PostRepository) Get(id uint) (post entity.Post, err error) {
	result := r.db.Preload("User").First(&post, id)
	return post, result.Error
}

func (r *PostRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Post{}, id)
	return result.Error
}

func (r *PostRepository) Update(id uint, post entity.Post) error {
	result := r.db.Model(&entity.Post{}).Where("id = ?", id).Updates(post)
	return result.Error
}
