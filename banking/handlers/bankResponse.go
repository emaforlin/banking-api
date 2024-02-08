package handlers

import "github.com/gin-gonic/gin"

type baseResponse struct {
	Message string `json: "message"`
}

func response(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, &baseResponse{
		Message: message,
	})
}
