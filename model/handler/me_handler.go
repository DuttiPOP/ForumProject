package handler

import (
	"ForumProject/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getMyProfile(c *gin.Context) {
	userID, err := h.GetUserID(c)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.services.IUserService.Get(userID)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) getMyPosts(c *gin.Context) {
	userID, err := h.GetUserID(c)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	posts, err := h.services.IUserService.GetAllPosts(userID)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, posts)
}
func (h *Handler) updateMyProfile(c *gin.Context) {
	userID, err := h.GetUserID(c)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	var updateDTO dto.UserUpdate
	err = c.BindJSON(&updateDTO)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.services.IUserService.Update(userID, updateDTO)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
