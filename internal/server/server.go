package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jwt-go/internal/database"
	"jwt-go/internal/handler"
	"jwt-go/internal/repository"
	"jwt-go/internal/service"
	"jwt-go/util"
)

func Run(cfg util.Config) {
	app := gin.Default()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "page not found",
		})
	})

	_repos := repository.NewRepos(database.Connect(cfg)) // initialize database
	_service := service.NewService(_repos)               // initialize service
	handlers := handler.NewHandler(_service)             // initialize handlers

	handlers.InitRoutes(app)

	err := app.Run()
	if err != nil {
		logrus.Fatalln("server down")
	}

	logrus.Println("server is running...")
}
