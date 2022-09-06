package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetMarks(c *gin.Context) {
	marks, err := h.automaticYoutube.Marks.GetMarks(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Get marks")

	c.JSON(http.StatusOK, marks)
}

func (h *Handler) AddMark(c *gin.Context) {
	var mark domain.MarkCreate
	if err := c.BindJSON(&mark); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Marks.AddMark(c, mark); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Add mark %s", mark.Title)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) UpdateMark(c *gin.Context) {
	markID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var mark domain.MarkCreate
	if err := c.BindJSON(&mark); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Marks.UpdateMark(c, markID, mark); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Update mark %s", markID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteMark(c *gin.Context) {
	markID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Marks.DeleteMark(c, markID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("Delete mark %s", markID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
