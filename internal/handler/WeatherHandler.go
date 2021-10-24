package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const URL = "https://api.weatherapi.com/v1/"
const API_KEY = "6d7078767a144e609be234341212310"

func (handler *Handler) Weather(c *gin.Context) {
	id, _ := getUserId(c)
	data, _ := handler.services.User.GetList(id)
	if data.Location.City == "" {
		Send(c, http.StatusInternalServerError, "Your location is unknown")
	}

	resp, err := http.Get(URL + "current.json?key=" + API_KEY + "&q=" + data.Location.City)
	if err != nil {
		Send(c, http.StatusBadRequest, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Send(c, http.StatusBadRequest, err.Error())
		return
	}

	var t map[string]interface{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		Send(c, http.StatusBadRequest, err.Error())
		return
	}

	Send(c, http.StatusOK, t["current"])
}

func (handler *Handler) Week(c *gin.Context) {
	id, _ := getUserId(c)
	data, _ := handler.services.User.GetList(id)

	if data.Location.City == "" {
		Send(c, http.StatusInternalServerError, "Your location is unknown")
	}

	resp, err := http.Get(URL + "forecast.json?key=" + API_KEY + "&q=" + data.Location.City + "&days=7")
	if err != nil {
		Send(c, http.StatusBadRequest, err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Send(c, http.StatusBadRequest, err.Error())
		return
	}

	var t map[string]interface{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		Send(c, http.StatusBadRequest, err.Error())
		return
	}

	Send(c, http.StatusOK, t)
}
