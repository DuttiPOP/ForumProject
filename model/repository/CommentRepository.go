package repository

import (
	"ForumProject/model/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment entity.Comment) (entity.Comment, error) {
	result := r.db.Create(&comment)
	if result.Error != nil {
		return entity.Comment{}, result.Error
	}
	return comment, nil
}

func (r *CommentRepository) Get(id uint) (comment entity.Comment, err error) {
	result := r.db.Preload("User").First(&comment, id)
	return comment, result.Error
}

func (r *CommentRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Comment{}, id)
	return result.Error
}

func (r *CommentRepository) Update(id uint, comment entity.Comment) error {
	result := r.db.Model(&entity.Comment{}).Where("id = ?", id).Updates(comment)
	return result.Error
}
