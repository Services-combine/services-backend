package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Manager interface {
	CreateToken(userID primitive.ObjectID, tokenTTL time.Duration) (string, *Payload, error)
	VerifyToken(accessToken string) (*Payload, error)
}
