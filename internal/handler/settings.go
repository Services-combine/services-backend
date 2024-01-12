package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initSettingsRoutes(inviting *gin.RouterGroup) {
	settings := inviting.Group("/settings")
	{
		settings.GET("/", h.getSettings)
		settings.PATCH("/", h.saveSettings)
	}
}

// @Summary		Get Settings
// @Tags			settings
// @Description	get settings
// @ModuleID		getSettings
// @Accept			json
// @Produce		json
// @Success		200		{object}	settings.Settings
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/settings [get]
func (h *Handler) getSettings(c *gin.Context) {
	settings, err := h.services.Settings.Get(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, settings)
}

type SaveSettingsRequest struct {
	CountInviting int `json:"count_inviting"`
	CountMailing  int `json:"count_mailing"`
}

// @Summary		Save Settings
// @Tags			settings
// @Description	save settings
// @ModuleID		saveSettings
// @Accept			json
// @Produce		json
// @Param			input	body		SaveSettingsRequest	true	"settings info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/settings [patch]
func (h *Handler) saveSettings(c *gin.Context) {
	var req SaveSettingsRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Settings.Save(c, NewSaveSettingsInput(req)); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
