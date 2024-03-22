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
	if err = c.BindJSON(&input); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.validate.Struct(input); err != nil {
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
	if err = c.BindJSON(&updateDTO); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	if err = h.services.IPostService.Update(userID, postID, updateDTO); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, http.NoBody)

}
