package repository

import (
	"context"
	domain_settings "github.com/b0shka/services/internal/domain/settings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ServiceInviting = "inviting"
)

type SettingsRepo struct {
	db *mongo.Collection
}

func NewSettingsRepo(db *mongo.Database) *SettingsRepo {
	return &SettingsRepo{db: db.Collection(settingsCollection)}
}

func (r *SettingsRepo) Get(ctx context.Context, service string) (domain_settings.Settings, error) {
	var settings domain_settings.Settings

	err := r.db.FindOne(ctx, bson.M{"service": service}).Decode(&settings)
	return settings, err
}

type SaveSettingsParams struct {
	Service       string `json:"service"`
	CountInviting int    `json:"countInviting"`
	CountMailing  int    `json:"countMailing"`
}

func (r *SettingsRepo) Save(ctx context.Context, arg SaveSettingsParams) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"service": arg.Service}, bson.M{"$set": bson.M{"countInviting": arg.CountInviting, "countMailing": arg.CountMailing}})
	return err
}
