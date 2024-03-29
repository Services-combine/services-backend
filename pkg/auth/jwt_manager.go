package auth

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/b0shka/services/internal/domain"
	"github.com/golang-jwt/jwt"
)

const (
	minSecretKeyLength = 32
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) (Manager, error) {
	if len(secretKey) < minSecretKeyLength {
		return nil, fmt.Errorf("invalid key length: must be at least %d characters", minSecretKeyLength)
	}

	return &JWTManager{secretKey: secretKey}, nil
}

func (m *JWTManager) CreateToken(userID primitive.ObjectID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(m.secretKey))

	return token, payload, err
}

func (m *JWTManager) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, domain.ErrInvalidToken
		}

		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		var verr *jwt.ValidationError

		ok := errors.As(err, &verr)
		if ok && errors.Is(verr.Inner, domain.ErrExpiredToken) {
			return nil, domain.ErrExpiredToken
		}

		return nil, domain.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, domain.ErrInvalidToken
	}

	return payload, nil
}
