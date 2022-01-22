package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) MainPage(c *gin.Context) {
	folders, err := h.services.Folders.Get(c, "/")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"folders": folders,
	})
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

	if err := h.services.Folders.Create(c, folder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) OpenFolder(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folder, accounts, countAccounts, foldersMove, err := h.services.Folders.OpenFolder(c, folderID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"folder":        folder,
		"accounts":      accounts,
		"countAccounts": countAccounts,
		"foldersMove":   foldersMove,
		"path":          "",
	})
}

func (h *Handler) MoveFolder(c *gin.Context) {
	var folderMove domain.FolderMove

	if err := c.BindJSON(&folderMove); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.Move(c, folderID, folderMove.Path); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) RenameFolder(c *gin.Context) {
	var folderName domain.FolderRename

	if err := c.BindJSON(&folderName); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.Rename(c, folderID, folderName.Name); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeChatFolder(c *gin.Context) {
	var folderChat domain.FolderChat

	if err := c.BindJSON(&folderChat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeChat(c, folderID, folderChat.Chat); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeUsernamesFolder(c *gin.Context) {
	var folderUsernames domain.FolderUsernames

	if err := c.BindJSON(&folderUsernames); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeUsernames(c, folderID, folderUsernames.Usernames); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeMessageFolder(c *gin.Context) {
	var folderMessage domain.FolderMessage

	if err := c.BindJSON(&folderMessage); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeMessage(c, folderID, folderMessage.Message); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) ChangeGroupsFolder(c *gin.Context) {
	var folderGroups domain.FolderGroups

	if err := c.BindJSON(&folderGroups); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.ChangeGroups(c, folderID, folderGroups.Groups); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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

	if err := h.services.Folders.Delete(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) LaunchInviting(c *gin.Context) {
	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Folders.LaunchInviting(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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

	if err := h.services.Folders.LaunchMailingUsernames(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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

	if err := h.services.Folders.LaunchMailingGroups(c, folderID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
