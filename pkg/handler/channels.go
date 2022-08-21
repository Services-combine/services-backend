package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
)

func (h *Handler) AddChannel(c *gin.Context) {
	var channel domain.ChannelAdd

	formData, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	channel.ChannelId = formData.Value["channel_id"][0]
	channel.ApiKey = formData.Value["api_key"][0]
	tokenFile := formData.File["token_file"][0]
	tokenFilePath := os.Getenv("FOLDER_CHANNELS") + channel.ChannelId + ".json"

	if err := c.SaveUploadedFile(tokenFile, tokenFilePath); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка при скачивании токен файла")
		return
	}

	if err := h.channels.Add(c, channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
