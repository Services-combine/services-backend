package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
)

func (h *Handler) Inviting(c *gin.Context) {
	folders, err := h.services.GetFolders(c, "/")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	fmt.Println(folders)
	c.JSON(http.StatusOK, map[string]interface{}{
		"folders": folders,
	})
}

func (h *Handler) CreateFolder(c *gin.Context) {
	var folder domain.Folder
	if err := c.BindJSON(&folder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	folder.Inviting = false
	folder.Mailing_usernames = false
	folder.Mailing_groups = false

	if err := h.services.CreateFolder(c, folder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
