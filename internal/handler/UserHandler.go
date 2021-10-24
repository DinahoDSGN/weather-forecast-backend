package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt-go/internal/models"
	"net/http"
	"strconv"
)

func (handler *Handler) User(c *gin.Context) {
	id, _ := getUserId(c)
	data, err := handler.services.GetList(id)
	if err != nil {
		Send(c, http.StatusInternalServerError, err.Error())
		return
	}

	if data.Id == 0 {
		Send(c, http.StatusInternalServerError, "Id not found")
		return
	}

	Send(c, http.StatusOK, id)
}

func (handler *Handler) GetAll(c *gin.Context) {
	data, _ := handler.services.GetAll()

	if data == nil {
		Send(c, http.StatusInternalServerError, "User table is empty")
		return
	}

	Send(c, http.StatusOK, data)
}

func (handler *Handler) GetList(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))

	if paramId <= 0 {
		Send(c, http.StatusBadRequest, "Invalid id")
		return
	}

	data, _ := handler.services.GetList(paramId)
	if data.Id == 0 {
		Send(c, http.StatusInternalServerError, "Failed to find by id")
		return
	}

	Send(c, http.StatusOK, data)
}

func (handler *Handler) Update(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))
	var JSONInput models.User

	if paramId <= 0 {
		Send(c, http.StatusBadRequest, "Invalid id")
		return
	}

	if err := c.BindJSON(&JSONInput); err != nil {
		Send(c, http.StatusBadRequest, "Cannot parse JSON into structure")
		return
	}

	data, _ := handler.services.Update(paramId, JSONInput)
	if data.Id == 0 {
		Send(c, http.StatusInternalServerError, "Failed to update by id")
		return
	}

	Send(c, http.StatusOK, data)
}

func (handler *Handler) Delete(c *gin.Context) {
	paramId, _ := strconv.Atoi(c.Param("id"))

	if paramId <= 0 {
		Send(c, http.StatusBadRequest, "Invalid id")
		return
	}

	data, _ := handler.services.Delete(paramId)
	if data.Id == 0 {
		Send(c, http.StatusInternalServerError, "Failed to delete by id")
		return
	}

	Send(c, http.StatusOK, fmt.Sprintf("%d successfully deleted", data.Id))
}
