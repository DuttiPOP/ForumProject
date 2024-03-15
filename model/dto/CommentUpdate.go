package dto

type CommentUpdate struct {
	Content string `json:"content" validate:"required,min=1,max=500"`
}
