package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateAccount(c *gin.Context) {
	var accountCreate domain.Account

	if err := c.BindJSON(&accountCreate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	phoneNew := strings.Replace(accountCreate.Phone, "+", "", 1)
	phoneNew = strings.Replace(phoneNew, "-", "", -1)
	phoneNew = strings.Replace(phoneNew, " ", "", -1)
	accountCreate.Phone = phoneNew

	folderID, err := primitive.ObjectIDFromHex(c.Param("folderID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accountCreate.Folder = folderID

	if err := h.inviting.Accounts.Create(c, accountCreate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Infof("CreateAccount %s", phoneNew)
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
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
