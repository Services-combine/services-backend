package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetChannels(c *gin.Context) {
	channels, err := h.automaticYoutube.Channels.Get(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, channels)
}

func (h *Handler) AddChannel(c *gin.Context) {
	var channel domain.ChannelIdKey

	formData, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	channel.ChannelId = formData.Value["channel_id"][0]
	channel.ApiKey = formData.Value["api_key"][0]

	status, err := h.automaticYoutube.CheckingUniqueness(c, channel.ChannelId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if status {
		appToken := formData.File["token_file"][0]
		appTokenPath := os.Getenv("FOLDER_CHANNELS") + "app_token_" + channel.ChannelId + ".json"
		userTokenPath := os.Getenv("FOLDER_CHANNELS") + "user_token_" + channel.ChannelId + ".json"

		if err := c.SaveUploadedFile(appToken, appTokenPath); err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrByDownloadTokenFile.Error())
			h.logger.Error(err)
			return
		}

		_, err = h.GetClient(c, appTokenPath, userTokenPath)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrUnableCreateUserToken.Error())
			h.logger.Error(err)
			return
		}

		if err := h.automaticYoutube.Add(c, channel); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	} else {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrChannelIdNoUniqueness.Error())
	}
}

func (h *Handler) LaunchChannel(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.ChannelIdKey
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	appTokenPath := os.Getenv("FOLDER_CHANNELS") + "app_token_" + channel.ChannelId + ".json"
	userTokenPath := os.Getenv("FOLDER_CHANNELS") + "user_token_" + channel.ChannelId + ".json"

	_, err = h.GetClient(c, appTokenPath, userTokenPath)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrUnableCreateUserToken.Error())
		h.logger.Error(err)
		return
	}

	if err := h.automaticYoutube.Channels.Launch(c, channelID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) UpdateChannel(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.ChannelIdKey
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Channels.Update(c, channelID, channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteChannel(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.ChannelIdKey
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Channels.Delete(c, channelID, channel.ChannelId); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) EditChannel(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.ChannelEdit
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Channels.Edit(c, channelID, channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
