package dto

type UserOutput struct {
	ID       uint            `json:"id" validate:"required"`
	Username string          `json:"username" validate:"required"`
	Posts    []PostOutput    `json:"posts,omitempty" validate:"required,dive"`
	Comments []CommentOutput `json:"comments,omitempty" validate:"required,dive"`
}
