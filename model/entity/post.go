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
	User     *User     `gorm:"foreignkey:UserID"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewPost(input dto.PostInput, userID uint) *Post {
	return &Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}
}

func UpdatePost(updateDTO dto.PostUpdate, postID uint, userID uint) *Post {
	var post Post
	post.ID = postID
	post.UserID = userID
	if updateDTO.Title != "" {
		post.Title = updateDTO.Title
	}
	if updateDTO.Content != "" {
		post.Content = updateDTO.Content
	}
	return &post
}
