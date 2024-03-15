package dto

type UserUpdate struct {
	Username string `json:"username,omitempty" validate:"required,alphanum,min=3,max=255"`
	Email    string `json:"email,omitempty" validate:"required,email,max=255"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
}
