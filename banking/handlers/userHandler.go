package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	RegisterUser(c *gin.Context)
	QueryUser(c *gin.Context)
}
