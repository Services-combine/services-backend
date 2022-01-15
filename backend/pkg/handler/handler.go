package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("ui/templates/*.html")
	router.Static("/static", "./ui/static")

	//router.GET("/sign-in", h.signInLoad)
	router.POST("/sign-in", h.signIn)

	services := router.Group("/", h.userIdentity)
	{
		services.GET("/", h.Index)
		inviting := services.Group("/inviting")
		{
			inviting.GET("/", h.Inviting)
			inviting.POST("/create-folder", h.CreateFolder)
			inviting.GET("/:hash")
			inviting.POST("/:hash/create-folder")
		}
	}

	return router
}
