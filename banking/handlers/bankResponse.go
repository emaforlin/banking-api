package handlers

import "github.com/gin-gonic/gin"

type baseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func response(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, &baseResponse{
		Message: message,
		Data:    data,
	})
}
