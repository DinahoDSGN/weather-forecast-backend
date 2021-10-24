package handler

import "github.com/gin-gonic/gin"

func Send(context *gin.Context, status int, data interface{}) {
	context.JSON(status, gin.H{
		"data": data,
	})
}
