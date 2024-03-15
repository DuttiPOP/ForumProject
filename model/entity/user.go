package entity

import (
	"ForumProject/model/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"type:varchar(255);not null;unique" db:"username" validate:"required,alphanum,min=3,max=255"`
	Email    string    `gorm:"type:varchar(255);not null;unique" db:"email" validate:"required,email,max=255"`
	Password string    `gorm:"type:varchar(255);not null" db:"password" validate:"required,min=8"`
	Posts    []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewUser(input dto.SignUpInput) *User {
	return &User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}
}

func UpdateUser(input dto.UserUpdate) *User {
	var user User
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		user.Password = input.Password
	}
	return &user
}
