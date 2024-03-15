package handler

import (
	"ForumProject/model/constants"
	"ForumProject/model/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUserById(c *gin.Context) {
	id, err := utils.StrToUInt(c.Params.ByName(constants.AnotherUserIDKey))
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.services.IUserService.Get(id)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) getUserPosts(c *gin.Context) {
	userID, err := utils.StrToUInt(c.Params.ByName("user_id"))
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
