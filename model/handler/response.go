package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"error"`
}

type response struct {
	Status string `json:"status"`
}

func (h *Handler) sendErrorResponse(c *gin.Context, message string, code int) {
	c.AbortWithStatusJSON(code, errorResponse{message})
}
