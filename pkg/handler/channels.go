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
	var channel domain.ChannelAdd

	formData, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	channel.ChannelId = formData.Value["channel_id"][0]
	channel.ApiKey = formData.Value["api_key"][0]
	channel.Proxy = formData.Value["proxy"][0]
	channel.Mark = formData.Value["mark"][0]

	status, err := h.automaticYoutube.CheckingUniqueness(c, channel.ChannelId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if status {
		appToken := formData.File["app_token_file"][0]
		userToken := formData.File["user_token_file"][0]
		appTokenPath := os.Getenv("FOLDER_CHANNELS") + "app_token_" + channel.ChannelId + ".json"
		userTokenPath := os.Getenv("FOLDER_CHANNELS") + "user_token_" + channel.ChannelId + ".json"

		if err := c.SaveUploadedFile(appToken, appTokenPath); err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrByDownloadAppTokenFile.Error())
			h.logger.Error(err)
			return
		}

		if err := c.SaveUploadedFile(userToken, userTokenPath); err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrByDownloadUserTokenFile.Error())
			h.logger.Error(err)
			return
		}

		/*_, err = h.GetClient(c, appTokenPath, userTokenPath)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrUnableCreateUserToken.Error())
			h.logger.Error(err)
			return
		}*/

		if err := h.automaticYoutube.Add(c, channel); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Infof("AddChannel %s", channel.ChannelId)
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

	h.logger.Infof("LaunchChannel %s", c.Param("channelID"))
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

	h.logger.Infof("UpdateChannel %s", c.Param("channelID"))
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

	h.logger.Infof("DeleteChannel %s", c.Param("channelID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) EditComment(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.CommentEdit
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Channels.EditChannel(c, channelID, channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("EditComment %s", c.Param("channelID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) EditProxy(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.ProxyEdit
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Channels.EditProxy(c, channelID, channel.Proxy); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("EditProxy %s", c.Param("channelID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) EditMark(c *gin.Context) {
	channelID, err := primitive.ObjectIDFromHex(c.Param("channelID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var channel domain.MarkEdit
	if err := c.BindJSON(&channel); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	markID, err := primitive.ObjectIDFromHex(channel.Mark)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.automaticYoutube.Channels.EditMark(c, channelID, markID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("EditMark %s", c.Param("channelID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
