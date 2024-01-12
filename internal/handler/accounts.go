package handler

import (
	"github.com/b0shka/services/internal/domain"
	"github.com/b0shka/services/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (h *Handler) initAccountsRoutes(folders *gin.RouterGroup) {
	accounts := folders.Group("/:folderID")
	{
		accounts.POST("/create-account", h.createAccount)
		accounts.PATCH("/:accountID", h.updateAccount)
		accounts.DELETE("/:accountID", h.deleteAccount)
		accounts.GET("/generate-interval", h.generateInterval)
		accounts.GET("/check-block", h.checkBlock)
		accounts.GET("/join-group", h.joinGroup)
	}
}

// @Summary		Create Account
// @Tags			accounts
// @Description	create a new account
// @ModuleID		createAccount
// @Accept			json
// @Produce		json
// @Param			input	body		domain.Account	true	"account info"
// @Success		200		{string}	string	"ok
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/create-account [post]
func (h *Handler) createAccount(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	formData, err := c.MultipartForm()
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
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
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accountCreate.Api_id = apiId
	accountCreate.Api_hash = formData.Value["api_hash"][0]

	status, err := h.services.CheckingUniqueness(c, phone)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if status {
		sessionFile := formData.File["session_file"][0]
		sessionFilePath := os.Getenv("FOLDER_ACCOUNTS") + phone + ".session"

		if err := c.SaveUploadedFile(sessionFile, sessionFilePath); err != nil {
			newResponse(c, http.StatusBadRequest, domain.ErrByDownloadSessionFile.Error())
			logger.Error(err)
			return
		}

		if err := h.services.Accounts.Create(c, accountCreate); err != nil {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		logger.Infof("CreateAccount %s", phone)
		c.Status(http.StatusOK)
	} else {
		newResponse(c, http.StatusBadRequest, domain.ErrPhoneNoUniqueness.Error())
	}
}

// @Summary		Update Account
// @Tags			accounts
// @Description	update account
// @ModuleID		updateAccount
// @Accept			json
// @Produce		json
// @Param			input	body		domain.AccountUpdate	true	"account info"
// @Success		200		{string}	string	"ok
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/:accountID [patch]
func (h *Handler) updateAccount(c *gin.Context) {
	var accountUpdate domain.AccountUpdate

	if err := c.BindJSON(&accountUpdate); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountID, err := primitive.ObjectIDFromHex(c.Param("accountID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accountUpdate.ID = accountID

	folderObjectID, err := primitive.ObjectIDFromHex(accountUpdate.FolderID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accountUpdate.Folder = folderObjectID

	if err := h.services.Accounts.Update(c, accountUpdate); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("UpdateAccount %s", c.Param("accountID"))
	c.Status(http.StatusOK)
}

// @Summary		Delete Account
// @Tags			accounts
// @Description	delete account
// @ModuleID		deleteAccount
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/:accountID [delete]
func (h *Handler) deleteAccount(c *gin.Context) {
	accountID, err := primitive.ObjectIDFromHex(c.Param("accountID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Accounts.Delete(c, accountID); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("DeleteAccount %s", c.Param("accountID"))
	c.Status(http.StatusOK)
}

// @Summary		Generate Interval For Accounts
// @Tags			accounts
// @Description	generate interval for accounts
// @ModuleID		generateInterval
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/generate-interval [get]
func (h *Handler) generateInterval(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Accounts.GenerateInterval(c, folderID); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("GenerateInterval %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Check Block For Accounts
// @Tags			accounts
// @Description	check block for accounts
// @ModuleID		checkBlock
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/check-block [get]
func (h *Handler) checkBlock(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.services.Accounts.CheckBlock(c, folderID); err != nil {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}()

	logger.Infof("CheckBlock %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Join Group Accounts
// @Tags			accounts
// @Description	join group accounts
// @ModuleID		joinGroup
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/join-group [get]
func (h *Handler) joinGroup(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	go func() {
		if err := h.services.Accounts.JoinGroup(c, folderID); err != nil {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}()

	logger.Infof("JoinGroup %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}
