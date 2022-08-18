package repository

import (
	"context"

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

func (s *FoldersRepo) GetFoldersByPath(ctx context.Context, path string) ([]domain.FolderItem, error) {
	var folders []domain.FolderItem

	cur, err := s.db.Find(ctx, bson.M{"path": path})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &folders); err != nil {
		return nil, err
	}

	return folders, nil
}

func (s *FoldersRepo) GetFolders(ctx context.Context) ([]domain.Folder, error) {
	var folders []domain.Folder

	cur, err := s.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &folders); err != nil {
		return nil, err
	}

	return folders, err
}

func (s *FoldersRepo) GetCountAccounts(ctx context.Context, folderID primitive.ObjectID) (domain.AccountsCount, error) {
	var countAccounts domain.AccountsCount
	var accounts []domain.Account
	var cur *mongo.Cursor
	var err error

	if folderID == primitive.NilObjectID {
		cur, err = s.db.Database().Collection(accountsCollection).Find(ctx, bson.M{})
	} else {
		cur, err = s.db.Database().Collection(accountsCollection).Find(ctx, bson.M{"folder": folderID})
	}
	if err != nil {
		return domain.AccountsCount{}, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return domain.AccountsCount{}, err
	}

	for _, account := range accounts {
		countAccounts.All++
		if account.Status_block == "clean" {
			countAccounts.Clean++
		} else {
			countAccounts.Block++
		}
	}

	return countAccounts, nil
}

func (s *FoldersRepo) Create(ctx context.Context, folder domain.Folder) error {
	_, err := s.db.InsertOne(ctx, folder)
	return err
}

func (s *FoldersRepo) GetFolderById(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) {
	var folder domain.Folder
	var folderPath domain.Folder

	err := s.db.FindOne(ctx, bson.M{"_id": folderID}).Decode(&folder)
	if err != nil {
		return domain.Folder{}, err
	}

	if folder.Path != "/" {
		folderPathObjectID, err := primitive.ObjectIDFromHex(folder.Path)
		if err != nil {
			return domain.Folder{}, err
		}

		err = s.db.FindOne(ctx, bson.M{"_id": folderPathObjectID}).Decode(&folderPath)
		if err != nil {
			return domain.Folder{}, err
		}

		folder.NamePath = folderPath.Name
	} else {
		folder.NamePath = "/"
	}

	return folder, nil
}

func (s *FoldersRepo) GetAccountsByFolderID(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error) {
	var accounts []domain.Account

	cur, err := s.db.Database().Collection(accountsCollection).Find(ctx, bson.M{"folder": folderID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *FoldersRepo) Move(ctx context.Context, folderID primitive.ObjectID, path string) error {
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

func (s *FoldersRepo) LaunchInviting(ctx context.Context, folderID primitive.ObjectID) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"inviting": true}})
	if err != nil {
		return err
	}

	_, err = s.db.Database().Collection(accountsCollection).UpdateMany(ctx, bson.M{"folder": folderID}, bson.M{"$set": bson.M{"launch": true}})
	return err
}

func (s *FoldersRepo) LaunchMailingUsernames(ctx context.Context, folderID primitive.ObjectID) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"mailing_usernames": true}})
	if err != nil {
		return err
	}

	_, err = s.db.Database().Collection(accountsCollection).UpdateMany(ctx, bson.M{"folder": folderID}, bson.M{"$set": bson.M{"launch": true}})
	return err
}

func (s *FoldersRepo) LaunchMailingGroups(ctx context.Context, folderID primitive.ObjectID) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": folderID}, bson.M{"$set": bson.M{"mailing_groups": true}})
	if err != nil {
		return err
	}

	_, err = s.db.Database().Collection(accountsCollection).UpdateMany(ctx, bson.M{"folder": folderID}, bson.M{"$set": bson.M{"launch": true}})
	return err
}

func (s *FoldersRepo) GetSettings(ctx context.Context) (domain.Settings, error) {
	var settings domain.Settings

	err := s.db.Database().Collection(userDataCollection).FindOne(ctx, bson.M{"_id": "admin"}).Decode(&settings)
	return settings, err
}
