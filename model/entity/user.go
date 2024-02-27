package entity

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
}
