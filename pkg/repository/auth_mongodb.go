package repository

import (
	"context"

	combine "github.com/korpgoodness/services.git"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongoDB struct {
	db *mongo.Collection
}

func NewAuthMongoDB(db *mongo.Collection) *AuthMongoDB {
	return &AuthMongoDB{db: db}
}

func (s *AuthMongoDB) GetUser(ctx context.Context, username, password string) (combine.User, error) {
	var user combine.User
	filter := bson.M{"username": username, "password": password}

	err := s.db.FindOne(ctx, filter).Decode(&user)
	return user, err
}
