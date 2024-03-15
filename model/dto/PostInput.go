package dto

type PostInput struct {
	Title   string `json:"title" validate:"required,min=2,max=100"`
	Content string `json:"content" validate:"required,min=5"`
}
