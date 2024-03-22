package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"error"`
}

func (h *Handler) sendErrorResponse(c *gin.Context, message string, code int) {
	c.AbortWithStatusJSON(code, errorResponse{message})
}
