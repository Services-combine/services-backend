package service_test

import (
	"context"
	"errors"
	"github.com/b0shka/services/internal/config"
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	domain_user "github.com/b0shka/services/internal/domain/user"
	mock_repository "github.com/b0shka/services/internal/repository/mocks"
	"github.com/b0shka/services/internal/service"
	mock_worker "github.com/b0shka/services/internal/worker/mocks"
	"github.com/b0shka/services/pkg/auth"
	"github.com/b0shka/services/pkg/hash"
	"github.com/b0shka/services/pkg/identity"
	"github.com/b0shka/services/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
	"time"
)

var ErrInternalServerError = errors.New("test: internal server error")

func mockAuthService(t *testing.T) (
	*service.AuthService,
	*mock_repository.MockUsers,
) {
	repoCtl := gomock.NewController(t)
	defer repoCtl.Finish()

	workerCtl := gomock.NewController(t)
	defer workerCtl.Finish()

	repoUsers := mock_repository.NewMockUsers(repoCtl)
	worker := mock_worker.NewMockTaskDistributor(workerCtl)
	authService := service.NewAuthService(
		repoUsers,
		&hash.SHA256Hasher{},
		&auth.JWTManager{},
		config.AuthConfig{},
		worker,
	)

	return authService, repoUsers
}

//func TestUsersService_Login(t *testing.T) {
//	authService, userRepo := mockAuthService(t)
//
//	w := httptest.NewRecorder()
//	ctx, _ := gin.CreateTestContext(w)
//
//	userRepo.EXPECT().Get(ctx, gomock.Any())
//	userRepo.EXPECT().CreateSession(ctx, gomock.Any())
//
//	res, err := authService.Login(ctx, domain_auth.LoginInput{})
//	require.NoError(t, err)
//	require.IsType(t, domain_auth.LoginOutput{}, res)
//}

func TestUsersService_LoginErrGetUser(t *testing.T) {
	authService, userRepo := mockAuthService(t)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	userRepo.EXPECT().Get(ctx, gomock.Any()).
		Return(domain_user.User{}, ErrInternalServerError)

	res, err := authService.Login(ctx, domain_auth.LoginInput{})
	require.True(t, errors.Is(err, ErrInternalServerError))
	require.IsType(t, domain_auth.LoginOutput{}, res)
}

//func TestUsersService_SignInErrCreateSession(t *testing.T) {
//	authService, userRepo := mockAuthService(t)
//
//	w := httptest.NewRecorder()
//	ctx, _ := gin.CreateTestContext(w)
//
//	userRepo.EXPECT().Get(ctx, gomock.Any())
//	userRepo.EXPECT().CreateSession(ctx, gomock.Any()).Return(ErrInternalServerError)
//
//	res, err := authService.Login(ctx, domain_auth.LoginInput{})
//	require.True(t, errors.Is(err, ErrInternalServerError))
//	require.IsType(t, domain_auth.LoginOutput{}, res)
//}

func TestUsersService_RefreshToken(t *testing.T) {
	authService, userRepo := mockAuthService(t)

	duration := time.Minute
	userID := identity.NewIDGenerator().GenerateObjectID()

	symmetricKey, err := utils.RandomString(32)
	require.NoError(t, err)
	tokenManager, err := auth.NewPasetoManager(symmetricKey)
	require.NoError(t, err)

	token, payload, err := tokenManager.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	ctx := context.Background()
	userRepo.EXPECT().GetSession(ctx, gomock.Any())

	res, _ := authService.RefreshToken(ctx, domain_auth.RefreshTokenInput{
		RefreshToken: token,
	})
	// require.NoError(t, err)
	require.IsType(t, domain_auth.RefreshTokenOutput{}, res)
}
