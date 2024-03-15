package handler

import (
	"ForumProject/model/constants"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user_id")
		if user == nil {
			h.sendErrorResponse(c, "user not authenticated", http.StatusUnauthorized)
			return
		}
		c.Set(constants.UserIDKey, user)
		c.Next()
	}
}

func (h *Handler) GetUserID(c *gin.Context) (uint, error) {
	userIdValue, exists := c.Get(constants.UserIDKey)
	if !exists {
		return 0, errors.New("user ID does not exist in the context")
	}

	userId, ok := userIdValue.(int)
	if !ok {
		return 0, errors.New("cannot convert user ID to uint")
	}

	return uint(userId), nil
}
