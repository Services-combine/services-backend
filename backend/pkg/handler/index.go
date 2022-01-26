package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	//_, ok := c.Get(userCtx)
	//if !ok {
	//	newErrorResponse(c, http.StatusUnauthorized, "user not found")
	//	return
	//}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
