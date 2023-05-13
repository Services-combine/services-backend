package handler

import (
	"net/http"
	"strings"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	formData, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var accountCreate domain.Account
	accountCreate.Name = formData.Value["name"][0]
	accountCreate.Folder = folderID

	phone := strings.Replace(formData.Value["phone"][0], "+", "", 1)
	phone = strings.Replace(phone, "-", "", -1)
	phone = strings.Replace(phone, " ", "", -1)
	accountCreate.Phone = phone

	apiId, err := strconv.Atoi(formData.Value["api_id"][0])
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
    }
	accountCreate.Api_id = apiId
	accountCreate.Api_hash = formData.Value["api_hash"][0]

	status, err := h.inviting.CheckingUniqueness(c, phone)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if status {
		sessionFile := formData.File["session_file"][0]
		sessionFilePath := os.Getenv("FOLDER_ACCOUNTS") + phone + ".session"

		if err := c.SaveUploadedFile(sessionFile, sessionFilePath); err != nil {
			newErrorResponse(c, http.StatusBadRequest, domain.ErrByDownloadSessionFile.Error())
			h.logger.Error(err)
			return
		}

		if err := h.inviting.Accounts.Create(c, accountCreate); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		h.logger.Infof("CreateAccount %s", phone)
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	} else {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrPhoneNoUniqueness.Error())
	}
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	var accountUpdate domain.AccountUpdate

	if err := c.BindJSON(&accountUpdate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountID, err := primitive.ObjectIDFromHex(c.Param("accountID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accountUpdate.ID = accountID

	folderObjectID, err := primitive.ObjectIDFromHex(accountUpdate.FolderID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accountUpdate.Folder = folderObjectID

	if err := h.inviting.Accounts.Update(c, accountUpdate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("UpdateAccount %s", c.Param("accountID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	accountID, err := primitive.ObjectIDFromHex(c.Param("accountID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Accounts.Delete(c, accountID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("DeleteAccount %s", c.Param("accountID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) GenerateInterval(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Accounts.GenerateInterval(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("GenerateInterval %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) CheckBlock(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.inviting.Accounts.CheckBlock(c, folderID); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}()

	h.logger.Infof("CheckBlock %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) JoinGroup(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.inviting.Accounts.JoinGroup(c, folderID); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}()

	h.logger.Infof("JoinGroup %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
