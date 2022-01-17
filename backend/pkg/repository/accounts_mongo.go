package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountsRepo struct {
	db *mongo.Collection
}

func NewAccountsRepo(db *mongo.Database) *AccountsRepo {
	return &AccountsRepo{db: db.Collection(accountsCollection)}
}

func (s *AccountsRepo) Create(ctx context.Context, accountCreate domain.Account) error {
	_, err := s.db.InsertOne(ctx, accountCreate)
	return err
}

func (s *AccountsRepo) GetData(ctx context.Context, accountID primitive.ObjectID) (domain.Account, error) {
	var account domain.Account

	err := s.db.FindOne(ctx, bson.M{"_id": accountID}).Decode(&account)
	return account, err
}

func (s *AccountsRepo) GetFolderByID(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) {
	var folder domain.Folder

	err := s.db.FindOne(ctx, bson.M{"_id": folderID}).Decode(&folder)
	return folder, err
}

func (s *AccountsRepo) Delete(ctx context.Context, accountID primitive.ObjectID) error {
	_, err := s.db.DeleteOne(ctx, bson.M{"_id": accountID})
	return err
}
