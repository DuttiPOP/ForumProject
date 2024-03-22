package handler

import (
	"ForumProject/model/constants"
	"ForumProject/model/new_errors"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(constants.UserIDKey)
		if user == nil {
			h.sendErrorResponse(c, new_errors.ErrUserNotAuthenticated, http.StatusUnauthorized)
			return
		}
		c.Set(constants.UserIDKey, user)
		c.Next()
	}
}

func (h *Handler) GetUserID(c *gin.Context) (uint, error) {
	userIdValue, exists := c.Get(constants.UserIDKey)
	if !exists {
		return 0, errors.New(new_errors.ErrUserIDDoesNotExist)
	}

	userId, ok := userIdValue.(int)
	if !ok {
		return 0, errors.New(new_errors.ErrUserIDCanNotBeConverted)
	}

	return uint(userId), nil
}
