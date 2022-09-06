package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
)

func (h *Handler) GetMarks(c *gin.Context) {
	marks, err := h.automaticYoutube.Marks.GetMarks(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, marks)
}

func (h *Handler) UpdateMark(c *gin.Context) {
	var marks domain.MarkGet
	if err := c.BindJSON(&marks); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Marks.UpdateMark(c, marks); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Save marks")

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
