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

func (r *SettingsRepo) GetMarks(ctx context.Context) ([]domain.Mark, error) {
	var marks domain.Marks

	err := r.db.FindOne(ctx, bson.M{"service": service_automatic_youtube}).Decode(&marks)
	if err != nil {
		return nil, err
	} else {
		return marks.Marks, nil
	}
}

func (r *SettingsRepo) SaveMarks(ctx context.Context, marks []domain.Mark) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"service": service_automatic_youtube}, bson.M{"$set": bson.M{"marks": marks}})
	return err
}

func (r *SettingsRepo) DeleteMark(ctx context.Context, mark domain.Mark) error {
	return nil
}
