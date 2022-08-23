package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
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

	dataUser, err := h.authorization.Authorization.CheckUser(c, userIdObject)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrNoAccessThisPage.Error())
		return
	}

	c.JSON(http.StatusOK, dataUser)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", domain.ErrHeaderAuthorizedIsEmpty
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", domain.ErrInvalidHeaderAuthorized
	}

	if len(headerParts[1]) == 0 {
		return "", domain.ErrTokenIsEmpty
	}

	return h.authorization.Authorization.ParseToken(headerParts[1])
}
