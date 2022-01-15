package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvitingMongoDB struct {
	db *mongo.Collection
}

func NewInvitingMongoDB(db *mongo.Database) *InvitingMongoDB {
	return &InvitingMongoDB{db: db.Collection(foldersCollection)}
}

func (s *InvitingMongoDB) GetFolders(ctx context.Context, path string) ([]domain.Folder, error) {
	var folders []domain.Folder

	cur, err := s.db.Find(ctx, bson.M{"path": path})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &folders); err != nil {
		return nil, err
	}

	return folders, nil
}

func (s *InvitingMongoDB) CreateFolder(ctx context.Context, folder domain.Folder) error {
	_, err := s.db.InsertOne(ctx, folder)
	if err != nil {
		return err
	}

	return nil
}
