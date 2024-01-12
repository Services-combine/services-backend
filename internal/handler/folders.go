package handler

import (
	"github.com/b0shka/services/internal/domain"
	"github.com/b0shka/services/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func (h *Handler) initFoldersRoutes(inviting *gin.RouterGroup) {
	folders := inviting.Group("/folders")
	{
		h.initAccountsRoutes(folders)

		folders.GET("/", h.getFolders)
		folders.POST("/create", h.createFolder)
		folders.GET("/:folderID", h.getFolderById)
		folders.DELETE("/:folderID", h.deleteFolder)
		folders.GET("/:folderID/folders-move", h.getFoldersMove)
		folders.POST("/:folderID/create-folder", h.createFolder)
		folders.POST("/:folderID/move", h.moveFolder)
		folders.POST("/:folderID/rename", h.renameFolder)
		folders.POST("/:folderID/change-chat", h.changeChat)
		folders.POST("/:folderID/change-usernames", h.changeUsernames)
		folders.POST("/:folderID/change-message", h.changeMessage)
		folders.POST("/:folderID/change-groups", h.changeGroups)
		folders.GET("/:folderID/launch-inviting", h.launchInviting)
		folders.GET("/:folderID/launch-mailing-usernames", h.launchMailingUsernames)
		folders.GET("/:folderID/launch-mailing-groups", h.launchMailingGroups)
	}
}

type GetFoldersResponse struct {
	Folders       []domain.FolderItem  `json:"folders"`
	CountAccounts domain.AccountsCount `json:"count_accounts"`
}

// @Summary		Get Folders
// @Tags			folders
// @Description	get folders
// @ModuleID		getFolders
// @Accept			json
// @Produce		json
// @Success		200		{object}	GetFoldersResponse
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders [get]
func (h *Handler) getFolders(c *gin.Context) {
	res, err := h.services.Folders.GetFolders(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, NewGetFoldersResponse(res))
}

// @Summary		Create Folders
// @Tags			folders
// @Description	create folders
// @ModuleID		createFolder
// @Accept			json
// @Produce		json
// @Param			input	body		domain.Folder	true	"folder info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/create [post]
func (h *Handler) createFolder(c *gin.Context) {
	var folder domain.Folder
	var path string

	if c.Param("folderID") == "" {
		path = "/"
	} else {
		path = c.Param("folderID")
	}

	if err := c.BindJSON(&folder); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folder.Path = path
	folder.Inviting = false
	folder.Mailing_usernames = false
	folder.Mailing_groups = false

	if err := h.services.Folders.Create(c, folder); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("CreateFolder %s", path)
	c.Status(http.StatusOK)
}

type GetFolderResponse struct {
	Folder        domain.Folder            `json:"folder"`
	Accounts      []domain.Account         `json:"accounts"`
	AccountsMove  []domain.AccountDataMove `json:"accounts_move"`
	Folders       []domain.FolderItem      `json:"folders"`
	CountAccounts domain.AccountsCount     `json:"count_accounts"`
	PathHash      []domain.AccountDataMove `json:"path_hash"`
}

// @Summary		Get Folder
// @Tags			folders
// @Description	get folders
// @ModuleID		getFolderById
// @Accept			json
// @Produce		json
// @Success		200		{object}	GetFolderResponse
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID [get]
func (h *Handler) getFolderById(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Folders.GetAllDataFolderById(c, folderID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, NewGetFolderResponse(res))
}

// @Summary		Get Folders For Move
// @Tags			folders
// @Description	get folders for move
// @ModuleID		getFoldersMove
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]domain.AccountDataMove
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/folders-move [get]
func (h *Handler) getFoldersMove(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderData, err := h.services.Folders.GetFoldersMove(c, folderID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, folderData)
}

// @Summary		Move Folder
// @Tags			folders
// @Description	move folder
// @ModuleID		moveFolder
// @Accept			json
// @Produce		json
// @Param			input	body		domain.FolderMove	true	"folder move info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/move [post]
func (h *Handler) moveFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderMove domain.FolderMove
	if err := c.BindJSON(&folderMove); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.Move(c, folderID, folderMove.Path); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("MoveFolder %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Rename Folder
// @Tags			folders
// @Description	rename folder
// @ModuleID		renameFolder
// @Accept			json
// @Produce		json
// @Param			input	body		domain.FolderRename	true	"folder rename info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/rename [post]
func (h *Handler) renameFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderName domain.FolderRename
	if err := c.BindJSON(&folderName); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.Rename(c, folderID, folderName.Name); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("RenameFolder %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Change Chat Folder
// @Tags			folders
// @Description	change chat folder
// @ModuleID		changeChat
// @Accept			json
// @Produce		json
// @Param			input	body		domain.FolderChat	true	"folder chat info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/change-chat [post]
func (h *Handler) changeChat(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderChat domain.FolderChat
	if err := c.BindJSON(&folderChat); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeChat(c, folderID, folderChat.Chat); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("ChangeChat %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Change Usernames Folder
// @Tags			folders
// @Description	change usernames folder
// @ModuleID		changeUsernames
// @Accept			json
// @Produce		json
// @Param			input	body		domain.FolderUsernames	true	"folder usernames info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/change-usernames [post]
func (h *Handler) changeUsernames(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderUsernames domain.FolderUsernames
	if err := c.BindJSON(&folderUsernames); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeUsernames(c, folderID, folderUsernames.Usernames); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("ChangeUsernames %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Change Message Folder
// @Tags			folders
// @Description	change message folder
// @ModuleID		changeMessage
// @Accept			json
// @Produce		json
// @Param			input	body		domain.FolderMessage	true	"folder message info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/change-message [post]
func (h *Handler) changeMessage(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderMessage domain.FolderMessage
	if err := c.BindJSON(&folderMessage); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeMessage(c, folderID, folderMessage.Message); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("ChangeMessage %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Change Groups Folder
// @Tags			folders
// @Description	change groups folder
// @ModuleID		changeGroups
// @Accept			json
// @Produce		json
// @Param			input	body		domain.FolderGroups	true	"folder groups info"
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/change-groups [post]
func (h *Handler) changeGroups(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderGroups domain.FolderGroups
	if err := c.BindJSON(&folderGroups); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeGroups(c, folderID, folderGroups.Groups); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("ChangeGroups %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

type DeleteFolderResponse struct {
	Path string `json:"path"`
}

// @Summary		Delete Folder
// @Tags			folders
// @Description	delete folder
// @ModuleID		deleteFolder
// @Accept			json
// @Produce		json
// @Success		200		{object}	DeleteFolderResponse
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID [delete]
func (h *Handler) deleteFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folder, err := h.services.Folders.GetFolderById(c, folderID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.Delete(c, folderID); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("DeleteFolder %s", c.Param("folderID"))
	c.JSON(http.StatusOK, DeleteFolderResponse{folder.Path})
}

// @Summary		Launch Inviting Folder
// @Tags			folders
// @Description	launch inviting folder
// @ModuleID		launchInviting
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/launch-inviting [get]
func (h *Handler) launchInviting(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.LaunchInviting(c, folderID); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("LaunchInviting %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Launch Mailing Usernames Folder
// @Tags			folders
// @Description	launch mailing usernames folder
// @ModuleID		launchMailingUsernames
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/launch-mailing-usernames [get]
func (h *Handler) launchMailingUsernames(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.LaunchMailingUsernames(c, folderID); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("LaunchMailingUsernames %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}

// @Summary		Launch Mailing Groups Folder
// @Tags			folders
// @Description	launch mailing groups folder
// @ModuleID		launchMailingGroups
// @Accept			json
// @Produce		json
// @Success		200		{string}	string	"ok"
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/inviting/folders/:folderID/launch-mailing-groups [get]
func (h *Handler) launchMailingGroups(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.LaunchMailingGroups(c, folderID); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logger.Infof("LaunchMailingGroups %s", c.Param("folderID"))
	c.Status(http.StatusOK)
}
