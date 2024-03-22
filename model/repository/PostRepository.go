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

func (r *PostRepository) Create(post entity.Post) (entity.Post, error) {
	result := r.db.Create(&post)
	if result.Error != nil {
		return entity.Post{}, result.Error
	}
	return post, nil
}

func (r *PostRepository) Get(id uint) (post entity.Post, err error) {
	result := r.db.Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		First(&post, id)
	return post, result.Error
}

func (r *PostRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Post{}, id)
	return result.Error
}

func (r *PostRepository) Update(post entity.Post) error {
	result := r.db.Model(&entity.Post{}).Where("id = ?", post.ID).Updates(post)
	return result.Error
}
