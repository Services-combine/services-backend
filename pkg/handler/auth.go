package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	combine "github.com/korpgoodness/services.git"
)

func (h *Handler) singIn(c *gin.Context) {
	var input combine.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(c, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
