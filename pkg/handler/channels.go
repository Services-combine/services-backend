package handler

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
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
		tokenFile := formData.File["token_file"][0]
		tokenFilePath := os.Getenv("FOLDER_CHANNELS") + "app_token_" + channel.ChannelId + ".json"
		if err := c.SaveUploadedFile(tokenFile, tokenFilePath); err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrByDownloadTokenFile.Error())
			return
		}

		// create user_token.json

		b, err := ioutil.ReadFile(tokenFilePath)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Unable to read client secret file")
			return
		}

		config, err := google.ConfigFromJSON(b, youtube.YoutubeForceSslScope, youtube.YoutubeUploadScope)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Unable to parse client secret file to config")
			return
		}
		config.RedirectURL = os.Getenv("URL_LISTEN_OAUTH_CODE")

		client := getClient(c, config)
		_, err = youtube.New(client)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Error creating YouTube client")
			return
		}

		//

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
