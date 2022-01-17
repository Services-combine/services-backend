package repository

import (
	"context"
	"fmt"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *FoldersRepo) GetData(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) {
	var folder domain.Folder

	err := s.db.FindOne(ctx, bson.M{"_id": folderID}).Decode(&folder)
	return folder, err
}

func (s *FoldersRepo) Move(ctx context.Context, folderID primitive.ObjectID, path string) error {
	fmt.Println(folderID, path)
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"path": path}})
	return err
}

func (s *FoldersRepo) Rename(ctx context.Context, folderID primitive.ObjectID, name string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"name": name}})
	return err
}

func (s *FoldersRepo) ChangeChat(ctx context.Context, folderID primitive.ObjectID, chat string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"chat": chat}})
	return err
}

func (s *FoldersRepo) ChangeUsernames(ctx context.Context, folderID primitive.ObjectID, usernames []string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"usernames": usernames}})
	return err
}

func (s *FoldersRepo) ChangeMessage(ctx context.Context, folderID primitive.ObjectID, message string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"message": message}})
	return err
}

func (s *FoldersRepo) ChangeGroups(ctx context.Context, folderID primitive.ObjectID, groups []string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"groups": groups}})
	return err
}

func (s *FoldersRepo) Delete(ctx context.Context, folderID primitive.ObjectID) error {
	_, err := s.db.DeleteOne(ctx, bson.M{"_id": folderID})
	return err
}
