package handler

import (
	"fmt"
	"github.com/b0shka/services/pkg/auth"
	"github.com/b0shka/services/pkg/identity"
	"github.com/b0shka/services/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func addAuthorizationHeader(
	t *testing.T,
	request *http.Request,
	tokenManager auth.Manager,
	authorizationType string,
	userID primitive.ObjectID,
	duration time.Duration,
) {
	token, payload, err := tokenManager.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestHandler_userIDentity(t *testing.T) {
	userID := identity.NewIDGenerator().GenerateObjectID()

	testTable := []struct {
		name         string
		setupAuth    func(t *testing.T, request *http.Request, tokenManager auth.Manager)
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			setupAuth: func(t *testing.T, request *http.Request, tokenManager auth.Manager) {
				addAuthorizationHeader(t, request, tokenManager, authorizationTypeBearer, userID, time.Minute)
			},
			statusCode:   200,
			responseBody: "",
		},
		{
			name: "no authorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenManager auth.Manager) {
			},
			statusCode:   401,
			responseBody: `{"message":"empty authorization header"}`,
		},
		{
			name: "unsupported authorization",
			setupAuth: func(t *testing.T, request *http.Request, tokenManager auth.Manager) {
				addAuthorizationHeader(t, request, tokenManager, "unsupported", userID, time.Minute)
			},
			statusCode:   401,
			responseBody: fmt.Sprintf(`{"message":"unsupported authorization type: %s"}`, "unsupported"),
		},
		{
			name: "invalid authorization format",
			setupAuth: func(t *testing.T, request *http.Request, tokenManager auth.Manager) {
				addAuthorizationHeader(t, request, tokenManager, "", userID, time.Minute)
			},
			statusCode:   401,
			responseBody: `{"message":"invalid authorization header format"}`,
		},
		{
			name: "expired token",
			setupAuth: func(t *testing.T, request *http.Request, tokenManager auth.Manager) {
				addAuthorizationHeader(t, request, tokenManager, authorizationTypeBearer, userID, -time.Minute)
			},
			statusCode:   401,
			responseBody: `{"message":"token has expired"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			symmetricKey, err := utils.RandomString(32)
			require.NoError(t, err)

			tokenManager, err := auth.NewPasetoManager(symmetricKey)
			require.NoError(t, err)

			router := gin.Default()

			router.GET(
				"/identity",
				userIdentity(tokenManager),
				func(c *gin.Context) {
					c.Status(http.StatusOK)
				},
			)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/identity", nil)

			testCase.setupAuth(t, req, tokenManager)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testCase.statusCode, recorder.Code)
			assert.Equal(t, testCase.responseBody, recorder.Body.String())
		})
	}
}

func TestGetUserPayload(t *testing.T) {
	userID := identity.NewIDGenerator().GenerateObjectID()

	payload, err := auth.NewPayload(userID, time.Minute)
	require.NoError(t, err)

	normalContext := &gin.Context{}
	normalContext.Set(userCtx, payload)

	key, err := utils.RandomString(10)
	require.NoError(t, err)

	invalidContext := &gin.Context{}
	invalidContext.Set(userCtx, key)

	tests := []struct {
		name      string
		ctx       *gin.Context
		payload   *auth.Payload
		shouldErr bool
	}{
		{
			name:      "ok",
			ctx:       normalContext,
			payload:   payload,
			shouldErr: false,
		},
		// {
		// 	name:      "empty user id",
		// 	ctx:       &gin.Context{},
		// 	shouldErr: true,
		// },
		// {
		// 	name:      "invalid payload",
		// 	ctx:       invalidContext,
		// 	shouldErr: true,
		// },
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			payload, err := getUserPayload(testCase.ctx)

			if testCase.shouldErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.payload, payload)
		})
	}
}
