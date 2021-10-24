package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"jwt-go/util"
	"log"
	"net/http"
)

func (handler *Handler) Weather(c *gin.Context) {
	config, err := util.LoadConfig(".") // initialize config
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	id, _ := getUserId(c)
	data, _ := handler.services.User.GetList(id)
	if data.Location.City == "" {
		Send(c, http.StatusInternalServerError, "Your location is unknown")
	}

	resp, err := http.Get(config.URL + "current.json?key=" + config.API_KEY + "&q=" + data.Location.City)
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
	config, err := util.LoadConfig(".") // initialize config
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	id, _ := getUserId(c)
	data, _ := handler.services.User.GetList(id)

	if data.Location.City == "" {
		Send(c, http.StatusInternalServerError, "Your location is unknown")
	}

	resp, err := http.Get(config.URL + "forecast.json?key=" + config.API_KEY + "&q=" + data.Location.City + "&days=7")
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
