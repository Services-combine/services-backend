package repository

import (
	"context"
	"errors"
	"time"

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

func (r *AuthRepo) SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{
		"$set": bson.M{
			"session": session,
		}})

	return err
}

func (r *AuthRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.User, error) {
	var user domain.User
	if err := r.db.FindOne(ctx, bson.M{
		"session.refreshToken": refreshToken,
		"session.expiresAt":    bson.M{"$gt": time.Now()},
	}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err

	}

	return user, nil
}

func (r *AuthRepo) RemoveRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"session.refreshToken": refreshToken}, bson.M{
		"$set": bson.M{
			"session.refreshToken": "",
		}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFound
		}
		return err
	}

	return nil
}
