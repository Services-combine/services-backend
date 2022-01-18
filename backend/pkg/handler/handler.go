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
			inviting.GET("/", h.MainPage)
			inviting.POST("/create-folder", h.CreateFolder)
			inviting.GET("/:folderID", h.OpenFolder)
			inviting.POST("/:folderID/create-folder", h.CreateFolder)

			inviting.POST("/:folderID/move", h.MoveFolder)
			inviting.POST("/:folderID/rename", h.RenameFolder)
			inviting.POST("/:folderID/change-chat", h.ChangeChatFolder)
			inviting.POST("/:folderID/change-usernames", h.ChangeUsernamesFolder)
			inviting.POST("/:folderID/change-message", h.ChangeMessageFolder)
			inviting.POST("/:folderID/change-groups", h.ChangeGroupsFolder)
			inviting.POST("/:folderID/create-account", h.CreateAccount)
			inviting.GET("/:folderID/generate-interval", h.GenerateInterval)
			inviting.GET("/:folderID/check-block")
			inviting.GET("/:folderID/delete", h.DeleteFolder)
			inviting.GET("/:folderID/launch-inviting", h.LaunchInviting)
			inviting.GET("/:folderID/launch-mailing-usernames", h.LaunchMailingUsernames)
			inviting.GET("/:folderID/launch-mailing-groups", h.LaunchMailingGroups)

			inviting.POST("/:folderID/:accountID", h.UpdateAccount)
			inviting.GET("/:folderID/:accountID", h.OpenAccount)
			inviting.GET("/:folderID/:accountID/delete", h.DeleteAccount)
			inviting.GET("/:folderID/:accountID/login")
			inviting.GET("/:folderID/:accountID/send-code")
			inviting.POST("/:folderID/:accountID/parsing-api")
			inviting.POST("/:folderID/:accountID/verify")
		}
	}

	return router
}
