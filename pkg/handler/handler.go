package handler

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/korpgoodness/service.git/pkg/logging"
	"github.com/korpgoodness/service.git/pkg/service"
)

type Handler struct {
	services *service.Service
	logger   logging.Logger
}

func NewHandler(services *service.Service) *Handler {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type,access-control-allow-origin, access-control-allow-headers,authorization,my-custom-header"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	api := router.Group("/api")
	{
		api.POST("/login", h.Login)
		api.GET("/refresh", h.Refresh)

		services := api.Group("/user", h.userIdentity)
		{
			services.GET("/", h.ServicesPage)
			services.POST("/save-settings", h.SaveSettings)
			services.GET("/logout", h.Logout)

			inviting := services.Group("/inviting")
			{
				inviting.GET("/", h.MainPage)
				inviting.POST("/create-folder", h.CreateFolder)
				inviting.POST("/:folderID", h.OpenFolder)
				inviting.GET("/:folderID/folders-move", h.GetFoldersMove)
				inviting.POST("/:folderID/create-folder", h.CreateFolder)

				inviting.POST("/:folderID/move", h.MoveFolder)
				inviting.POST("/:folderID/rename", h.RenameFolder)
				inviting.POST("/:folderID/change-chat", h.ChangeChatFolder)
				inviting.POST("/:folderID/change-usernames", h.ChangeUsernamesFolder)
				inviting.POST("/:folderID/change-message", h.ChangeMessageFolder)
				inviting.POST("/:folderID/change-groups", h.ChangeGroupsFolder)
				inviting.POST("/:folderID/create-account", h.CreateAccount)
				inviting.GET("/:folderID/generate-interval", h.GenerateInterval)
				inviting.GET("/:folderID/check-block", h.CheckBlock)
				inviting.GET("/:folderID/delete", h.DeleteFolder)
				inviting.GET("/:folderID/launch-inviting", h.LaunchInviting)
				inviting.GET("/:folderID/launch-mailing-usernames", h.LaunchMailingUsernames)
				inviting.GET("/:folderID/launch-mailing-groups", h.LaunchMailingGroups)

				inviting.POST("/:folderID/:accountID", h.UpdateAccount)
				inviting.GET("/:folderID/:accountID", h.OpenAccount)
				inviting.GET("/:folderID/:accountID/delete", h.DeleteAccount)
				inviting.GET("/:folderID/:accountID/login-api", h.LoginApi)
				inviting.POST("/:folderID/:accountID/parsing-api", h.ParsingApi)
				inviting.GET("/:folderID/:accountID/get-code-session", h.GetCodeSession)
				inviting.POST("/:folderID/:accountID/create-session", h.CreateSession)
			}
		}
	}

	return router
}
