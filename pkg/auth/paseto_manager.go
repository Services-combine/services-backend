package auth

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/b0shka/services/internal/domain"
	"github.com/o1egl/paseto"
)

type PasetoManager struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoManager(symmetricKey string) (Manager, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key length: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	return &PasetoManager{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (m *PasetoManager) CreateToken(userID primitive.ObjectID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", nil, err
	}

	token, err := m.paseto.Encrypt(m.symmetricKey, payload, nil)

	return token, payload, err
}

func (m *PasetoManager) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
