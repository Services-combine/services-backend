package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongoDB struct {
	db *mongo.Collection
}

func NewAuthMongoDB(db *mongo.Database) *AuthMongoDB {
	return &AuthMongoDB{db: db.Collection(usersCollection)}
}

func (s *AuthMongoDB) GetUser(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User
	filter := bson.M{"username": username, "password": password}

	err := s.db.FindOne(ctx, filter).Decode(&user)
	return user, err
}
