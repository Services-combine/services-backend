package handler

import (
	"net/http"

	"github.com/b0shka/services/docs"
	"github.com/b0shka/services/internal/config"
	"github.com/b0shka/services/internal/service"
	"github.com/b0shka/services/pkg/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.Manager
}

func NewHandler(
	services *service.Services,
	tokenManager auth.Manager,
) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTP.Host
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := router.Group("/api/v1")
	{
		h.initAuthRoutes(api)
		h.initInvitingRoutes(api)
	}

	return router
}
