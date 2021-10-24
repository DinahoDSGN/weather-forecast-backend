package handler

import "github.com/gin-gonic/gin"

func (handler *Handler) Init(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
