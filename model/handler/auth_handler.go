package handler

import (
	"ForumProject/model/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input dto.SignUpInput
	if err := c.BindJSON(&input); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(input); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.services.IUserService.Create(input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) signIn(c *gin.Context) {
	var input dto.SignInInput
	if err := c.BindJSON(&input); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(input); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.services.IUserService.Authenticate(input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user)
	if err = session.Save(); err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
