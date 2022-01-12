package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/services.git/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sing-in", h.singIn)

	return router
}
