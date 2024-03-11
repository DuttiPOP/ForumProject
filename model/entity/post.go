package entity

import (
	"ForumProject/model/dto"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(255);not null" json:"title" db:"title" validate:"required,max=255"`
	Content  string    `gorm:"type:text;not null" json:"content" db:"content" validate:"required"`
	UserID   uint      `gorm:"not null" json:"user_id" db:"user_id"`
	User     User      `gorm:"foreignkey:UserID"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewPost(input dto.NewPostInput) *Post {
	return &Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
	}
}
