package handler

import (
	"ForumProject/model/constants"
	"ForumProject/model/dto"
	"ForumProject/model/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPost(c *gin.Context) {
	userID, err := h.GetUserID(c)
	var input dto.PostInput
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
	postID, err := h.services.IPostService.Create(userID, input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, postID)
}

func (h *Handler) getPostById(c *gin.Context) {
	id, err := utils.StrToUInt(c.Params.ByName(constants.PostIDKey))
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := h.services.IPostService.Get(id)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) updatePost(c *gin.Context) {
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
	var updateDTO dto.PostUpdate
	err = c.BindJSON(&updateDTO)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.services.IPostService.Update(userID, postID, updateDTO)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})

}
