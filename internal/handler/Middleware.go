package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const userCtx = "userId"

func (handler *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		Send(c, http.StatusUnauthorized, "Empty Authorization header")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		Send(c, http.StatusUnauthorized, "Invalid Authorization header")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if len(headerParts[1]) == 0 {
		Send(c, http.StatusUnauthorized, "Token is empty")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := handler.services.Auth.ParseToken(headerParts[1])
	if err != nil {
		Send(c, http.StatusUnauthorized, err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if userId == 0 {
		Send(c, http.StatusInternalServerError, "Invalid access token")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set(userCtx, int(userId))
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	if id == 0 {
		return 0, errors.New("id not found")
	}

	return idInt, nil
}
