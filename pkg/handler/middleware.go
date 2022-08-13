package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	userID, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userIdObject, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dataUser, err := h.services.Authorization.CheckUser(c, userIdObject)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Нет доступа к этой странице")
		return
	}

	c.JSON(http.StatusOK, dataUser)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("Пустой заголовок Authorized")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("НЕ валидный заголовок Authorized")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("Токен пустой")
	}

	return h.services.Authorization.ParseToken(headerParts[1])
}
