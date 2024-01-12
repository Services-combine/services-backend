package handler

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"

	"github.com/b0shka/services/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/refresh", h.refreshToken)
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	SessionID    primitive.ObjectID `json:"session_id"`
	RefreshToken string             `json:"refresh_token"`
	AccessToken  string             `json:"access_token"`
}

// @Summary		User Login
// @Tags			auth
// @Description	user login
// @ModuleID		login
// @Accept			json
// @Produce		json
// @Param			input	body		LoginRequest	true	"login info"
// @Success		200		{object}	LoginResponse
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInput.Error())

		return
	}

	res, err := h.services.Auth.Login(c, NewLoginInput(req))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, NewLoginResponse(res))
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// @Summary		User Refresh Token
// @Tags			auth
// @Description	user refresh token
// @ModuleID		refreshToken
// @Accept			json
// @Produce		json
// @Param			input	body		RefreshTokenRequest	true	"refresh info"
// @Success		200		{object}	RefreshTokenResponse
// @Failure		400,401	{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router			/auth/refresh [post]
func (h *Handler) refreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInput.Error())

		return
	}

	res, err := h.services.Auth.RefreshToken(c, NewRefreshTokenInput(req))
	if err != nil {
		if errors.Is(err, domain.ErrSessionBlocked) ||
			errors.Is(err, domain.ErrIncorrectSessionUser) ||
			errors.Is(err, domain.ErrMismatchedSession) ||
			errors.Is(err, domain.ErrExpiredToken) ||
			errors.Is(err, domain.ErrInvalidToken) {
			newResponse(c, http.StatusUnauthorized, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, NewRefreshTokenResponse(res))
}
