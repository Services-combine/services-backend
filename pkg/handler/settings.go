package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
)

func (h *Handler) GetSettings(c *gin.Context) {
	settings, err := h.settings.Settings.GetSettings(c)
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

	if err := h.settings.Settings.SaveSettings(c, dataSettings); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Save new settings")

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) GetMarks(c *gin.Context) {
	marks, err := h.settings.Settings.GetMarks(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, marks)
}

func (h *Handler) SaveMarks(c *gin.Context) {
	var marks []domain.Mark
	if err := c.BindJSON(&marks); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.settings.SaveMarks(c, marks); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Save marks")

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteMark(c *gin.Context) {
	var mark domain.Mark
	if err := c.BindJSON(&mark); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.settings.DeleteMark(c, mark); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Delete mark")

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
