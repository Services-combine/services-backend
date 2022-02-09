package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	accessTokenTTL  = 30 * time.Minute
	refreshTokenTTL = 30 * 24 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	return &AuthService{repo: repo}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (userData, error) {
	user, err := s.repo.GetUser(ctx, username, generatePasswordHash(password))
	if err != nil {
		return userData{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (userData, error) {
	_, err := s.ParseToken(refreshToken)
	if err != nil {
		return userData{}, err
	}

	user, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return userData{}, err
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	err := s.repo.RemoveRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) CreateSession(ctx context.Context, userId primitive.ObjectID) (userData, error) {
	var (
		res userData
		err error
	)

	res.UserID = userId.Hex()
	res.AccessToken, err = NewJWT(userId.Hex(), accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = NewJWT(userId.Hex(), refreshTokenTTL)
	if err != nil {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(refreshTokenTTL),
	}

	err = s.repo.SetSession(ctx, userId, session)
	return res, err
}

func NewJWT(userId string, tokenTTL time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
