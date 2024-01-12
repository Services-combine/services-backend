package service

import (
	"context"
	"github.com/b0shka/services/internal/config"
	"github.com/b0shka/services/internal/domain"
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	"github.com/b0shka/services/internal/repository"
	"github.com/b0shka/services/internal/worker"
	"github.com/b0shka/services/pkg/auth"
	"github.com/b0shka/services/pkg/hash"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	formatTimeLayout = "Jan _2, 2006 15:04:05 (MST)"
)

type AuthService struct {
	repo            repository.Users
	hasher          hash.Hasher
	tokenManager    auth.Manager
	authConfig      config.AuthConfig
	taskDistributor worker.TaskDistributor
}

func NewAuthService(
	repo repository.Users,
	hasher hash.Hasher,
	tokenManager auth.Manager,
	authConfig config.AuthConfig,
	taskDistributor worker.TaskDistributor,
) *AuthService {
	return &AuthService{
		repo:            repo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		authConfig:      authConfig,
		taskDistributor: taskDistributor,
	}
}

func (s *AuthService) Login(ctx *gin.Context, inp domain_auth.LoginInput) (domain_auth.LoginOutput, error) {
	passwordHash, err := s.hasher.HashCode(inp.Password)
	if err != nil {
		return domain_auth.LoginOutput{}, err
	}

	userParams := repository.GetUserParams{
		Email:    inp.Email,
		Password: passwordHash,
	}
	user, err := s.repo.Get(ctx, userParams)
	if err != nil {
		return domain_auth.LoginOutput{}, err
	}

	tokens, err := s.createSession(ctx, user.ID)
	if err != nil {
		return domain_auth.LoginOutput{}, err
	}

	taskPayload := &worker.PayloadSendLoginNotification{
		Email:     inp.Email,
		UserAgent: ctx.Request.UserAgent(),
		ClientIP:  ctx.ClientIP(),
		Time:      time.Now().Format(formatTimeLayout),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(worker.QueueDefault),
	}

	err = s.taskDistributor.DistributeTaskSendLoginNotification(ctx, taskPayload, opts...)
	if err != nil {
		return domain_auth.LoginOutput{}, err
	}

	return tokens, nil
}

func (s *AuthService) createSession(ctx *gin.Context, id primitive.ObjectID) (domain_auth.LoginOutput, error) {
	refreshToken, refreshPayload, err := s.tokenManager.CreateToken(
		id,
		s.authConfig.JWT.RefreshTokenTTL,
	)
	if err != nil {
		return domain_auth.LoginOutput{}, err
	}

	accessToken, _, err := s.tokenManager.CreateToken(
		id,
		s.authConfig.JWT.AccessTokenTTL,
	)
	if err != nil {
		return domain_auth.LoginOutput{}, err
	}

	res := domain_auth.NewLoginOutput(
		refreshPayload.ID,
		refreshToken,
		accessToken,
	)

	sessionParams := repository.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       id,
		RefreshToken: res.RefreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIP:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiresAt,
	}

	if err := s.repo.CreateSession(ctx, sessionParams); err != nil {
		return domain_auth.LoginOutput{}, err
	}

	return res, nil
}

func (s *AuthService) RefreshToken(
	ctx context.Context,
	inp domain_auth.RefreshTokenInput,
) (domain_auth.RefreshTokenOutput, error) {
	var res domain_auth.RefreshTokenOutput

	refreshPayload, err := s.tokenManager.VerifyToken(inp.RefreshToken)
	if err != nil {
		return domain_auth.RefreshTokenOutput{}, err
	}

	session, err := s.repo.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		return domain_auth.RefreshTokenOutput{}, err
	}

	if session.IsBlocked {
		return domain_auth.RefreshTokenOutput{}, domain.ErrSessionBlocked
	}

	if refreshPayload.UserID != session.UserID {
		return domain_auth.RefreshTokenOutput{}, domain.ErrIncorrectSessionUser
	}

	if inp.RefreshToken != session.RefreshToken {
		return domain_auth.RefreshTokenOutput{}, domain.ErrMismatchedSession
	}

	if time.Now().After(session.ExpiresAt) {
		return domain_auth.RefreshTokenOutput{}, domain.ErrExpiredToken
	}

	accessToken, _, err := s.tokenManager.CreateToken(
		refreshPayload.UserID,
		s.authConfig.JWT.RefreshTokenTTL,
	)
	if err != nil {
		return domain_auth.RefreshTokenOutput{}, err
	}

	res.AccessToken = accessToken

	return res, nil
}
