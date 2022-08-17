package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/korpgoodness/service.git/internal/domain"
)

func (h *Handler) Login(c *gin.Context) {
	var inp domain.User
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	res, err := h.authorization.Authorization.Login(c, inp.Username, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Login user %s", res.UserID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": res.AccessToken,
		"id":          res.UserID,
	})
}