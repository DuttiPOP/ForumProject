package dto

type PostOutput struct {
	ID       uint            `json:"id" validate:"required"`
	Title    string          `json:"title" validate:"required,min=1,max=100"`
	Content  string          `json:"content" validate:"required,min=1,max=500"`
	User     *UserOutput     `json:"user,omitempty" validate:"required"`
	Comments []CommentOutput `json:"comments"`
}
