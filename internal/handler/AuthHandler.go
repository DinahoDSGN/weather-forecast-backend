package handler

import (
	"github.com/gin-gonic/gin"
	"jwt-go/internal/models"
	"net/http"
)

type SignInForm struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	IPAddress string `json:"IPAddress"`
}

func (handler *Handler) SignIn(c *gin.Context) {
	var JSONInput SignInForm
	if err := c.BindJSON(&JSONInput); err != nil {
		Send(c, http.StatusBadRequest, "Cannot parse JSON into structure")
		return
	}

	token, err := handler.services.Auth.GenerateToken(JSONInput.Username, JSONInput.Password)
	if err != nil {
		Send(c, http.StatusInternalServerError, err.Error())
		return
	}

	Send(c, http.StatusOK, token)
}

func (handler *Handler) SignUp(c *gin.Context) {
	var JSONInput models.User
	if err := c.BindJSON(&JSONInput); err != nil {
		Send(c, http.StatusBadRequest, "Cannot parse JSON into structure")
		return
	}

	data, _ := handler.services.Auth.SignUp(JSONInput)
	if data.Id == 0 {
		Send(c, http.StatusBadRequest, "Invalid fields")
		return
	}

	Send(c, http.StatusOK, data.Username)
}
