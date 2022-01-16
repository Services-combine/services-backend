package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FoldersRepo struct {
	db *mongo.Collection
}

func NewFoldersRepo(db *mongo.Database) *FoldersRepo {
	return &FoldersRepo{db: db.Collection(foldersCollection)}
}

func (s *FoldersRepo) Get(ctx context.Context, path string) ([]domain.Folder, error) {
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

func (s *FoldersRepo) Create(ctx context.Context, folder domain.Folder) error {
	_, err := s.db.InsertOne(ctx, folder)
	return err
}

func (s *FoldersRepo) GetData(ctx context.Context, hash string) (domain.Folder, error) {
	var folder domain.Folder

	err := s.db.FindOne(ctx, bson.M{"hash": hash}).Decode(&folder)
	return folder, err
}

func (s *FoldersRepo) Move(ctx context.Context, hash, path string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"hash": hash}, bson.M{"$set": bson.M{"path": path}})
	return err
}

func (s *FoldersRepo) Rename(ctx context.Context, hash, name string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"hash": hash}, bson.M{"$set": bson.M{"name": name}})
	return err
}

func (s *FoldersRepo) ChangeChat(ctx context.Context, hash, chat string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"hash": hash}, bson.M{"$set": bson.M{"chat": chat}})
	return err
}

func (s *FoldersRepo) ChangeUsernames(ctx context.Context, hash string, usernames []string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"hash": hash}, bson.M{"$set": bson.M{"usernames": usernames}})
	return err
}

func (s *FoldersRepo) ChangeMessage(ctx context.Context, hash, message string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"hash": hash}, bson.M{"$set": bson.M{"message": message}})
	return err
}

func (s *FoldersRepo) ChangeGroups(ctx context.Context, hash string, groups []string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"hash": hash}, bson.M{"$set": bson.M{"groups": groups}})
	return err
}

func (s *FoldersRepo) Delete(ctx context.Context, hash string) error {
	_, err := s.db.DeleteOne(ctx, bson.M{"hash": hash})
	return err
}
