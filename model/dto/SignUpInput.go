package dto

type SignUpInput struct {
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}
