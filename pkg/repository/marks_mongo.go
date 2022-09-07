package repository

import (
	"context"
	"errors"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MarksRepo struct {
	db *mongo.Collection
}

func NewMarksRepo(db *mongo.Database) *MarksRepo {
	return &MarksRepo{db: db.Collection(marksCollection)}
}

func (r *MarksRepo) GetMarks(ctx context.Context) ([]domain.MarkGet, error) {
	var marks []domain.MarkGet

	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &marks); err != nil {
		return nil, err
	}

	return marks, nil
}

func (r *MarksRepo) AddMark(ctx context.Context, mark domain.MarkCreate) error {
	_, err := r.db.InsertOne(ctx, mark)
	return err
}

func (r *MarksRepo) UpdateMark(ctx context.Context, markID primitive.ObjectID, mark domain.MarkCreate) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": markID}, bson.M{"$set": mark})
	return err
}

func (r *MarksRepo) CheckMarkToDelete(ctx context.Context, markID primitive.ObjectID) (bool, error) {
	var channel domain.ChannelGet

	err := r.db.Database().Collection(channelsCollection).FindOne(ctx, bson.M{"mark": markID}).Decode(&channel)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (r *MarksRepo) DeleteMark(ctx context.Context, markID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": markID})
	return err
}
