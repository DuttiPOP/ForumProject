package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;unique" json:"username" db:"username"`
	Email    string `gorm:"type:varchar(255);not null;unique" json:"email" db:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password,omitempty" db:"password"`
}
