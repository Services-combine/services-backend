package auth

import (
	"github.com/b0shka/services/pkg/identity"
	"testing"
	"time"

	"github.com/b0shka/services/internal/domain"
	"github.com/b0shka/services/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestAuthJWT_NewJWTManager(t *testing.T) {
	validKey, err := utils.RandomString(32)
	require.NoError(t, err)

	invalidKey, err := utils.RandomString(31)
	require.NoError(t, err)

	tests := []struct {
		name      string
		key       string
		shouldErr bool
	}{
		{
			name:      "ok",
			key:       validKey,
			shouldErr: false,
		},
		{
			name:      "invalid key length",
			key:       invalidKey,
			shouldErr: true,
		},
		{
			name:      "invalid key length",
			key:       "",
			shouldErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			manager, err := NewJWTManager(testCase.key)

			if testCase.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.IsType(t, &JWTManager{}, manager)
			}
		})
	}
}

func TestAuthJWT_CreateTokenAndVerify(t *testing.T) {
	secretKey, err := utils.RandomString(32)
	require.NoError(t, err)
	manager, err := NewJWTManager(secretKey)
	require.NoError(t, err)

	userID := identity.NewIDGenerator().GenerateObjectID()

	duration := time.Minute
	testPayload, err := NewPayload(userID, duration)
	require.NoError(t, err)

	token, payload, err := manager.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	tokenExpired, payload, err := manager.CreateToken(userID, -duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	jwtTokenWithNoneSigning := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	tokenWithNoneSigning, err := jwtTokenWithNoneSigning.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	tests := []struct {
		name          string
		payload       *Payload
		token         string
		shouldErr     bool
		expectedError error
	}{
		{
			name:      "ok",
			payload:   testPayload,
			token:     token,
			shouldErr: false,
		},
		{
			name:          "invalid token",
			payload:       testPayload,
			token:         "",
			shouldErr:     true,
			expectedError: domain.ErrInvalidToken,
		},
		{
			name:          "invalid token",
			payload:       testPayload,
			token:         "token",
			shouldErr:     true,
			expectedError: domain.ErrInvalidToken,
		},
		{
			name:          "invalid token",
			payload:       testPayload,
			token:         tokenWithNoneSigning,
			shouldErr:     true,
			expectedError: domain.ErrInvalidToken,
		},
		{
			name:          "expired token",
			payload:       testPayload,
			token:         tokenExpired,
			shouldErr:     true,
			expectedError: domain.ErrExpiredToken,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			payload, err = manager.VerifyToken(testCase.token)

			if testCase.shouldErr {
				require.Error(t, err)
				require.Equal(t, err, testCase.expectedError)
				require.Nil(t, payload)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, payload)

				require.NotZero(t, payload.ID)
				require.Equal(t, testCase.payload.UserID, payload.UserID)
				require.WithinDuration(t, testCase.payload.IssuedAt, payload.IssuedAt, time.Second)
				require.WithinDuration(t, testCase.payload.ExpiresAt, payload.ExpiresAt, time.Second)
			}
		})
	}
}
