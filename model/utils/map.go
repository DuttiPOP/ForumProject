package utils

import (
	"ForumProject/model/dto"
	"ForumProject/model/entity"
)

func MapToPostDTO(post entity.Post) dto.PostOutput {
	postDTO := dto.PostOutput{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}

	if post.User != nil {
		userDTO := MapToUserDTO(*post.User)
		postDTO.User = &userDTO
	}
	if post.Comments != nil {
		postDTO.Comments = make([]dto.CommentOutput, len(post.Comments))
		for i, comment := range post.Comments {
			postDTO.Comments[i] = MapToCommentDTO(comment)
		}
	}
	return postDTO
}

func MapToUserDTO(user entity.User) dto.UserOutput {
	userDTO := dto.UserOutput{
		ID:       user.ID,
		Username: user.Username,
	}
	//if user.Posts != nil {
	//	userDTO.posts := make([]post_dto.PostOutput, len(user.Posts))
	//	for i, comment := range user.Posts {
	//		posts[i] = MapToCommentDTO(comment)
	//	}
	//}
	return userDTO

}

func MapToCommentDTO(comment entity.Comment) dto.CommentOutput {
	return dto.CommentOutput{
		ID:      comment.ID,
		Content: comment.Content,
		User:    MapToUserDTO(comment.User),
	}
}
