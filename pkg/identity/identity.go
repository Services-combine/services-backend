package identity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Generator interface {
	GenerateUUID() uuid.UUID
	GenerateObjectID() primitive.ObjectID
}

type IDGenerator struct{}

func NewIDGenerator() Generator {
	return &IDGenerator{}
}

func (g *IDGenerator) GenerateUUID() uuid.UUID {
	return uuid.New()
}

func (g *IDGenerator) GenerateObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}
