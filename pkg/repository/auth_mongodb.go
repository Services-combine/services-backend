package repository

import (
	"context"
	"errors"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	db *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) *AuthRepo {
	return &AuthRepo{db: db.Collection(usersCollection)}
}

func (r *AuthRepo) GetUser(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User

	if err := r.db.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}

func (r *AuthRepo) CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error) {
	var user domain.UserReduxData

	if err := r.db.FindOne(ctx, bson.M{"_id": userID}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserReduxData{}, domain.ErrUserNotFound
		}
		return domain.UserReduxData{}, err
	}

	return user, nil
}
