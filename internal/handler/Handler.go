package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jwt-go/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) InitRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/", handler.Init)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/signin", handler.SignIn)
		authRouter.POST("/signup", handler.SignUp)
	}

	apiRouter := router.Group("/api", handler.UserIdentity)

	userGroup := apiRouter.Group("/user")
	{
		userGroup.GET("/me", handler.User)
		userGroup.GET("/all", handler.GetAll)
		userGroup.GET("/:id", handler.GetList)
		userGroup.PUT("/:id", handler.Update)
		userGroup.DELETE("/:id", handler.Delete)
	}

	weatherGroup := apiRouter.Group("/weather")
	{
		weatherGroup.GET("/today", handler.Weather)
		weatherGroup.GET("/week", handler.Week)
	}

	logrus.Println("handlers is running...")

	return router
}
