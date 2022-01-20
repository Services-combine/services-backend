package repository

import (
	"context"
	"math/rand"
	"time"

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

func RandomInterval() uint8 {
	min := 15
	max := 40
	rand.Seed(time.Now().UnixNano())
	interval := min + rand.Intn(max-min+1)

	return uint8(interval)
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

	err := s.db.Database().Collection(foldersCollection).FindOne(ctx, bson.M{"_id": folderID}).Decode(&folder)
	return folder, err
}

func (s *AccountsRepo) UpdateAccount(ctx context.Context, account domain.AccountUpdate) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": account.ID}, bson.M{"$set": bson.M{"name": account.Name, "folder": account.Folder, "interval": account.Interval}})
	return err
}

func (s *AccountsRepo) Delete(ctx context.Context, accountID primitive.ObjectID) error {
	_, err := s.db.DeleteOne(ctx, bson.M{"_id": accountID})
	return err
}

func (s *AccountsRepo) GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error {
	var accounts []domain.Account

	cur, err := s.db.Find(ctx, bson.M{"folder": folderID})
	if err != nil {
		return err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return err
	}

	for _, account := range accounts {
		var accountU domain.AccountUpdate
		accountU.ID = account.ID
		accountU.Name = account.Name
		accountU.Folder = account.Folder
		accountU.Interval = RandomInterval()

		if err := s.UpdateAccount(ctx, accountU); err != nil {
			return err
		}
	}

	return nil
}

func (s *AccountsRepo) AddRandomHash(ctx context.Context, accountID primitive.ObjectID, randomHash string) error {
	_, err := s.db.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": bson.M{"random_hash": randomHash}})
	return err
}
