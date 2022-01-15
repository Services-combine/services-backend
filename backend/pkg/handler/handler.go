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
			inviting.GET("/:hash", h.OpenFolder)
			inviting.POST("/:hash/create-folder", h.CreateFolder)

			inviting.POST("/:hash/create-account")
			inviting.POST("/:hash/move-folder")
			inviting.POST("/:hash/rename-folder", h.RenameFolder)
			inviting.POST("/:hash/change-chat")
			inviting.POST("/:hash/add-usernames")
			inviting.POST("/:hash/add-messages")
			inviting.POST("/:hash/add-groups")
			inviting.GET("/:hash/generate-interval")
			inviting.GET("/:hash/check-block")
			inviting.GET("/:hash/delete-folder")
			inviting.GET("/:hash/launch-inviting")
			inviting.GET("/:hash/launch-mailing-usernames")
			inviting.GET("/:hash/launch-mailing-groups")

			inviting.GET("/:hash/:id")
			inviting.GET("/:hash/:id/delete-account")
			inviting.GET("/:hash/:id/login-account")
			inviting.GET("/:hash/:id/send-code-account")
			inviting.POST("/:hash/:id/parsing-api")
			inviting.POST("/:hash/:id/verify-account")
		}
	}

	return router
}
