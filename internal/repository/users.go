package repository

import (
	"context"
	"errors"
	"github.com/b0shka/services/internal/domain"
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	domain_user "github.com/b0shka/services/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UsersRepo struct {
	db *mongo.Collection
}

func NewUsersRepo(db *mongo.Database) *UsersRepo {
	return &UsersRepo{db: db.Collection(usersCollection)}
}

type CreateUserParams struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type GetUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UsersRepo) Get(ctx context.Context, arg GetUserParams) (domain_user.User, error) {
	var user domain_user.User

	if err := r.db.FindOne(ctx, bson.M{"email": arg.Email, "password": arg.Password}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain_user.User{}, domain.ErrUserNotFound
		}
		return domain_user.User{}, err
	}

	return user, nil
}

func (r *UsersRepo) Check(ctx context.Context, id primitive.ObjectID) (domain_user.UserReduxData, error) {
	var user domain_user.UserReduxData

	if err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain_user.UserReduxData{}, domain.ErrUserNotFound
		}
		return domain_user.UserReduxData{}, err
	}

	return user, nil
}

type CreateSessionParams struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	UserAgent    string             `json:"user_agent" bson:"user_agent"`
	ClientIP     string             `json:"client_ip" bson:"client_ip"`
	IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
	ExpiresAt    time.Time          `json:"expires_at" bson:"expires_at"`
}

func (r *UsersRepo) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := r.db.Database().Collection(sessionsCollection).InsertOne(ctx, arg)
	return err
}

func (r *UsersRepo) GetSession(ctx context.Context, id primitive.ObjectID) (domain_auth.Session, error) {
	var session domain_auth.Session

	if err := r.db.Database().Collection(sessionsCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&session); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain_auth.Session{}, domain.ErrSessionNotFound
		}
		return domain_auth.Session{}, err
	}

	return session, nil
}
