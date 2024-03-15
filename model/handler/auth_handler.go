package handler

import (
	"ForumProject/model/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input dto.SignUpInput
	err := c.BindJSON(&input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.validate.Struct(input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.services.IUserService.Create(input)
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func (h *Handler) signIn(c *gin.Context) {
	var input dto.SignInInput
	err := c.BindJSON(&input)
	if err != nil {
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
	err = session.Save()
	if err != nil {
		h.sendErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
