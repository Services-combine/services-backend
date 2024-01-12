package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) initInvitingRoutes(api *gin.RouterGroup) {
	inviting := api.Group("/inviting")
	inviting.Use(userIdentity(h.tokenManager))
	{
		h.initSettingsRoutes(inviting)
		h.initFoldersRoutes(inviting)
	}
}
