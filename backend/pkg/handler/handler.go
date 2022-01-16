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

	router.POST("/sign-in", h.signIn)

	services := router.Group("/", h.userIdentity)
	{
		services.GET("/", h.Index)
		inviting := services.Group("/inviting")
		{
			inviting.GET("/", h.Inviting)
			inviting.POST("/create", h.CreateFolder)
			inviting.GET("/:hash", h.OpenFolder)
			inviting.POST("/:hash/create", h.CreateFolder)

			inviting.POST("/:hash/create-account")
			inviting.POST("/:hash/move", h.MoveFolder)
			inviting.POST("/:hash/rename", h.RenameFolder)
			inviting.POST("/:hash/change-chat", h.ChangeChatFolder)
			inviting.POST("/:hash/change-usernames", h.ChangeUsernamesFolder)
			inviting.POST("/:hash/change-messages")
			inviting.POST("/:hash/change-groups", h.ChangeGroupsFolder)
			inviting.GET("/:hash/generate-interval")
			inviting.GET("/:hash/check-block")
			inviting.GET("/:hash/delete", h.DeleteFolder)
			inviting.GET("/:hash/launch-inviting")
			inviting.GET("/:hash/launch-mailing-usernames")
			inviting.GET("/:hash/launch-mailing-groups")

			inviting.GET("/:hash/:id")
			inviting.GET("/:hash/:id/delete")
			inviting.GET("/:hash/:id/login")
			inviting.GET("/:hash/:id/send-code")
			inviting.POST("/:hash/:id/parsing-api")
			inviting.POST("/:hash/:id/verify")
		}
	}

	return router
}
