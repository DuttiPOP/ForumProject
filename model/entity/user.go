package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;unique" json:"username" db:"username"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email" db:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty" db:"password"`
	Username string    `gorm:"type:varchar(255);not null;unique" db:"username" validate:"required,alphanum,min=3,max=255"`
	Email    string    `gorm:"type:varchar(255);not null;unique" db:"email" validate:"required,email,max=255"`
	Password string    `gorm:"type:varchar(255);not null" db:"password" validate:"required,min=8"`
	Posts    []Post    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
