package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/b0shka/services/internal/domain"
	"github.com/b0shka/services/pkg/identity"
)

type Payload struct {
	ID        primitive.ObjectID `json:"id"`
	UserID    primitive.ObjectID `json:"user_id"`
	IssuedAt  time.Time          `json:"issued_at"`
	ExpiresAt time.Time          `json:"expires_at"`
}

func NewPayload(userID primitive.ObjectID, duration time.Duration) (*Payload, error) {
	idGenerator := identity.NewIDGenerator()

	payload := &Payload{
		ID:        idGenerator.GenerateObjectID(),
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return domain.ErrExpiredToken
	}

	return nil
}
