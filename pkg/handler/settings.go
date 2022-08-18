package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
)

func (h *Handler) GetSettings(c *gin.Context) {
	settings, err := h.inviting.Settings.GetSettings(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, settings)
}

func (h *Handler) SaveSettings(c *gin.Context) {
	var dataSettings domain.Settings

	if err := c.BindJSON(&dataSettings); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Settings.SaveSettings(c, dataSettings); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Save new settings")

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
