package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChannelsRepo struct {
	db *mongo.Collection
}

func NewChannelsRepo(db *mongo.Database) *ChannelsRepo {
	return &ChannelsRepo{db: db.Collection(channelsCollection)}
}

func (r *ChannelsRepo) Add(ctx context.Context, channel domain.ChannelAdd) error {
	_, err := r.db.InsertOne(ctx, channel)
	return err
}
