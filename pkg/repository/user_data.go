package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ID = "admin"
)

type UserDataRepo struct {
	db *mongo.Collection
}

func NewUserDataRepo(db *mongo.Database) *UserDataRepo {
	return &UserDataRepo{db: db.Collection(userDataCollection)}
}

func (s *UserDataRepo) GetSettings(ctx context.Context) (domain.Settings, error) {
	var settings domain.Settings

	err := s.db.FindOne(ctx, bson.M{"_id": ID}).Decode(&settings)
	return settings, err
}

func (s *UserDataRepo) SaveSettings(ctx context.Context, dataSettings domain.Settings) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": ID}, bson.M{"$set": bson.M{"countInviting": dataSettings.CountInviting, "countMailing": dataSettings.CountMailing}})
	return err
}
