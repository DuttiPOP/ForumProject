package handler

import (
	"ForumProject/model/constants"
	"ForumProject/model/dto"
	"ForumProject/model/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCommentsByPostId(c *gin.Context) {
	postID, err := utils.StrToUInt(c.Params.ByName("post_id"))
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	comments, err := h.services.IPostService.GetCommentsByPostId(postID)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) createComment(c *gin.Context) {
	userID, err := h.GetUserID(c)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	postID, err := utils.StrToUInt(c.Params.ByName(constants.PostIDKey))
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	var input dto.CommentInput
	err = c.BindJSON(&input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.validate.Struct(input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	commentID, err := h.services.ICommentService.Create(userID, postID, input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, commentID)
}

func (h *Handler) updateComment(c *gin.Context) {

}
