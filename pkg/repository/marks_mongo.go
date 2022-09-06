package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *MarksRepo) UpdateMark(ctx context.Context, mark domain.MarkGet) error {
	//_, err := r.db.UpdateOne(ctx, bson.M{"service": service_automatic_youtube}, bson.M{"$set": bson.M{"marks": marks}})
	return nil
}
