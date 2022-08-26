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
	authorization    *service.AuthorizationService
	inviting         *service.InvitingService
	automaticYoutube *service.AutomaticYoutubeService
	logger           logging.Logger
}

func NewHandler(
	authorization *service.AuthorizationService,
	inviting *service.InvitingService,
	automaticYoutube *service.AutomaticYoutubeService,
) *Handler {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
	}

	return &Handler{
		authorization:    authorization,
		inviting:         inviting,
		automaticYoutube: automaticYoutube,
		logger:           logger,
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
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.Login)
			auth.GET("/get-me", h.userIdentity)

			user := auth.Group("/user")
			{
				inviting := user.Group("/inviting")
				{
					inviting.GET("/get-folders", h.GetFolders)
					inviting.GET("/get-settings", h.GetSettings)
					inviting.POST("/save-settings", h.SaveSettings)
					inviting.POST("/create-folder", h.CreateFolder)

					inviting.GET("/:folderID", h.GetFolderById)
					inviting.GET("/:folderID/folders-move", h.GetFoldersMove)
					inviting.POST("/:folderID/create-folder", h.CreateFolder)
					inviting.POST("/:folderID/move", h.MoveFolder)
					inviting.POST("/:folderID/rename", h.RenameFolder)
					inviting.POST("/:folderID/change-chat", h.ChangeChat)
					inviting.POST("/:folderID/change-usernames", h.ChangeUsernames)
					inviting.POST("/:folderID/change-message", h.ChangeMessage)
					inviting.POST("/:folderID/change-groups", h.ChangeGroups)
					inviting.POST("/:folderID/create-account", h.CreateAccount)
					inviting.GET("/:folderID/generate-interval", h.GenerateInterval)
					inviting.GET("/:folderID/check-block", h.CheckBlock)
					inviting.GET("/:folderID/delete", h.DeleteFolder)
					inviting.GET("/:folderID/launch-inviting", h.LaunchInviting)
					inviting.GET("/:folderID/launch-mailing-usernames", h.LaunchMailingUsernames)
					inviting.GET("/:folderID/launch-mailing-groups", h.LaunchMailingGroups)

					inviting.POST("/:folderID/:accountID", h.UpdateAccount)
					inviting.GET("/:folderID/:accountID/delete", h.DeleteAccount)
					inviting.GET("/:folderID/:accountID/login-api", h.LoginApi)
					inviting.POST("/:folderID/:accountID/parsing-api", h.ParsingApi)
					inviting.GET("/:folderID/:accountID/get-code-session", h.GetCodeSession)
					inviting.POST("/:folderID/:accountID/create-session", h.CreateSession)
				}

				channels := user.Group("/channels")
				{
					channels.GET("/", h.GetChannels)
					channels.POST("/add", h.AddChannel)
					channels.POST("/:channelID/launch", h.LaunchChannel)
					channels.POST("/:channelID/update", h.UpdateChannel)
					channels.POST("/:channelID/delete", h.DeleteChannel)
					channels.POST("/:channelID/edit-channel", h.EditChannel)
					channels.POST("/:channelID/edit-proxy", h.EditProxy)
				}
			}
		}
	}

	return router
}
