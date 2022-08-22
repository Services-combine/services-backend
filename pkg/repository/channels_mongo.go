package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *ChannelsRepo) GetChannels(ctx context.Context) ([]domain.ChannelGet, error) {
	var channels []domain.ChannelGet

	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &channels); err != nil {
		return nil, err
	}

	return channels, nil
}

func (r *ChannelsRepo) LaunchChannel(ctx context.Context, channelID primitive.ObjectID) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": channelID}, bson.M{"$set": bson.M{"launch": true}})
	return err
}

func (r *ChannelsRepo) UpdateChannel(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelAdd) error {
	_, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": channelID},
		bson.M{"$set": channel},
	)
	return err
}

func (r *ChannelsRepo) DeleteChannel(ctx context.Context, channelID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": channelID})
	return err
}

func (r *ChannelsRepo) EditChannel(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelEdit) error {
	_, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": channelID},
		bson.M{"$set": channel},
	)
	return err
}
