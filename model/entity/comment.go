package entity

import (
	"ForumProject/model/dto"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignkey:UserID"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	Post    Post   `gorm:"foreignkey:PostID"`
}

func NewComment(input dto.CommentInput) *Comment {
	return &Comment{
		Content: input.Content,
	}

}
