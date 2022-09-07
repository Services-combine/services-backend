package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetFolders(c *gin.Context) {
	dataPage, err := h.inviting.Folders.GetFolders(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("GetFolders")
	c.JSON(http.StatusOK, dataPage)
}

func (h *Handler) CreateFolder(c *gin.Context) {
	var folder domain.Folder
	var path string

	if c.Param("folderID") == "" {
		path = "/"
	} else {
		path = c.Param("folderID")
	}

	if err := c.BindJSON(&folder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folder.Path = path
	folder.Inviting = false
	folder.Mailing_usernames = false
	folder.Mailing_groups = false

	if err := h.inviting.Folders.Create(c, folder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("CreateFolder %s", path)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) GetFolderById(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderData, err := h.inviting.Folders.GetAllDataFolderById(c, folderID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("GetFolderById %s", c.Param("folderID"))
	c.JSON(http.StatusOK, folderData)
}

func (h *Handler) GetFoldersMove(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderData, err := h.inviting.Folders.GetFoldersMove(c, folderID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("GetFoldersMove %s", c.Param("folderID"))
	c.JSON(http.StatusOK, folderData)
}

func (h *Handler) MoveFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderMove domain.FolderMove
	if err := c.BindJSON(&folderMove); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.Move(c, folderID, folderMove.Path); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("MoveFolder %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) RenameFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderName domain.FolderRename
	if err := c.BindJSON(&folderName); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.Rename(c, folderID, folderName.Name); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("RenameFolder %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeChat(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderChat domain.FolderChat
	if err := c.BindJSON(&folderChat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.ChangeChat(c, folderID, folderChat.Chat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("ChangeChat %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeUsernames(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderUsernames domain.FolderUsernames
	if err := c.BindJSON(&folderUsernames); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.ChangeUsernames(c, folderID, folderUsernames.Usernames); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("ChangeUsernames %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeMessage(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderMessage domain.FolderMessage
	if err := c.BindJSON(&folderMessage); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.ChangeMessage(c, folderID, folderMessage.Message); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("ChangeMessage %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeGroups(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var folderGroups domain.FolderGroups
	if err := c.BindJSON(&folderGroups); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.ChangeGroups(c, folderID, folderGroups.Groups); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("ChangeGroups %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) DeleteFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folder, err := h.inviting.Folders.GetFolderById(c, folderID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.Delete(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("DeleteFolder %s", c.Param("folderID"))
	c.JSON(http.StatusOK, folder.Path)
}

func (h *Handler) LaunchInviting(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.LaunchInviting(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("LaunchInviting %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) LaunchMailingUsernames(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.LaunchMailingUsernames(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("LaunchMailingUsernames %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) LaunchMailingGroups(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.inviting.Folders.LaunchMailingGroups(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("LaunchMailingGroups %s", c.Param("folderID"))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
