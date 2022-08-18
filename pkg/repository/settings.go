package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	service_inviting = "inviting"
)

type UserDataRepo struct {
	db *mongo.Collection
}

func NewUserDataRepo(db *mongo.Database) *UserDataRepo {
	return &UserDataRepo{db: db.Collection(settingsCollection)}
}

func (s *UserDataRepo) GetSettings(ctx context.Context) (domain.Settings, error) {
	var settings domain.Settings

	err := s.db.FindOne(ctx, bson.M{"service": service_inviting}).Decode(&settings)
	return settings, err
}

func (s *UserDataRepo) SaveSettings(ctx context.Context, dataSettings domain.Settings) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"service": service_inviting}, bson.M{"$set": bson.M{"countInviting": dataSettings.CountInviting, "countMailing": dataSettings.CountMailing}})
	return err
}
