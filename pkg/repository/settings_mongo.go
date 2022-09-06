package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	service_inviting          = "inviting"
	service_automatic_youtube = "automatic-youtube"
)

type SettingsRepo struct {
	db *mongo.Collection
}

func NewSettingsRepo(db *mongo.Database) *SettingsRepo {
	return &SettingsRepo{db: db.Collection(settingsCollection)}
}

func (r *SettingsRepo) GetSettings(ctx context.Context) (domain.Settings, error) {
	var settings domain.Settings

	err := r.db.FindOne(ctx, bson.M{"service": service_inviting}).Decode(&settings)
	return settings, err
}

func (r *SettingsRepo) SaveSettings(ctx context.Context, dataSettings domain.Settings) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"service": service_inviting}, bson.M{"$set": bson.M{"countInviting": dataSettings.CountInviting, "countMailing": dataSettings.CountMailing}})
	return err
}
