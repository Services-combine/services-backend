package repository

import (
	"context"
	"fmt"
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	domain_user "github.com/b0shka/services/internal/domain/user"
	"github.com/b0shka/services/pkg/auth"
	"github.com/b0shka/services/pkg/hash"
	"github.com/b0shka/services/pkg/identity"
	"github.com/b0shka/services/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomSession(t *testing.T, user domain_user.User) domain_auth.Session {
	sessionID := identity.NewIDGenerator().GenerateObjectID()

	symmetricKey, err := utils.RandomString(32)
	require.NoError(t, err)

	tokenManager, err := auth.NewPasetoManager(symmetricKey)
	require.NoError(t, err)

	refreshToken, _, err := tokenManager.CreateToken(user.ID, time.Hour)
	require.NoError(t, err)

	userAgent, err := utils.RandomString(20)
	require.NoError(t, err)

	clientIP, err := utils.RandomInt(1, 255)
	require.NoError(t, err)

	arg := CreateSessionParams{
		ID:           sessionID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		ClientIP: fmt.Sprintf(
			"%d:%d:%d:%d",
			clientIP,
			clientIP,
			clientIP,
			clientIP,
		),
		IsBlocked: false,
		ExpiresAt: time.Now().Add(time.Hour),
	}

	err = testRepos.CreateSession(context.Background(), arg)
	require.NoError(t, err)

	return domain_auth.Session{
		ID:           sessionID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		ClientIP: fmt.Sprintf(
			"%d:%d:%d:%d",
			clientIP,
			clientIP,
			clientIP,
			clientIP,
		),
		IsBlocked: false,
		ExpiresAt: arg.ExpiresAt,
	}
}

func createRandomUser(t *testing.T) domain_user.User {
	email, err := utils.RandomString(7)
	require.NoError(t, err)

	salt, err := utils.RandomString(32)
	require.NoError(t, err)

	hasher, err := hash.NewSHA256Hasher(salt)
	require.NoError(t, err)

	passwordHash, err := hasher.HashCode(email)
	require.NoError(t, err)

	user := domain_user.User{
		ID:       identity.NewIDGenerator().GenerateObjectID(),
		Email:    fmt.Sprintf("%s@ya.ru", email),
		Password: passwordHash,
	}

	return user
}

func TestRepository_CreateSession(t *testing.T) {
	user := createRandomUser(t)
	createRandomSession(t, user)
}

func TestRepository_GetSession(t *testing.T) {
	user := createRandomUser(t)
	session1 := createRandomSession(t, user)
	session2, err := testRepos.GetSession(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, session1.ID, session2.ID)
	require.Equal(t, session1.UserID, session2.UserID)
	require.Equal(t, session1.RefreshToken, session2.RefreshToken)
	require.Equal(t, session1.UserAgent, session2.UserAgent)
	require.Equal(t, session1.ClientIP, session2.ClientIP)
	require.Equal(t, session1.IsBlocked, session2.IsBlocked)
	require.WithinDuration(t, session1.ExpiresAt, session2.ExpiresAt, time.Second)
}
