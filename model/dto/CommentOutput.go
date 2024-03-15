package dto

type CommentOutput struct {
	ID      uint       `json:"id" validate:"required"`
	PostID  uint       `json:"postId,omitempty" validate:"required"`
	Content string     `json:"content" validate:"required"`
	User    UserOutput `json:"user" validate:"required"`
}
