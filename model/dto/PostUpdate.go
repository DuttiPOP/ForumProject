package dto

type PostUpdate struct {
	Title   string `json:"title,omitempty" validate:"required,min=2,max=100"`
	Content string `json:"content,omitempty" validate:"required,min=5"`
}
