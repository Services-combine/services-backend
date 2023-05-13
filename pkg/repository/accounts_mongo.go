package repository

import (
	"context"
	"errors"
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

func (r *AccountsRepo) CheckingUniqueness(ctx context.Context, phone string) (bool, error) {
	var account domain.Account

	err := r.db.FindOne(ctx, bson.M{"phone": phone}).Decode(&account)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}

func (r *AccountsRepo) Create(ctx context.Context, accountCreate domain.Account) error {
	_, err := r.db.InsertOne(ctx, accountCreate)
	return err
}

func (r *AccountsRepo) GetById(ctx context.Context, accountID primitive.ObjectID) (domain.Account, error) {
	var account domain.Account

	err := r.db.FindOne(ctx, bson.M{"_id": accountID}).Decode(&account)
	return account, err
}

func (r *AccountsRepo) GetAccountsByFolderID(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error) {
	var accounts []domain.Account

	cur, err := r.db.Find(ctx, bson.M{"folder": folderID})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *AccountsRepo) Update(ctx context.Context, account domain.AccountUpdate) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": account.ID}, bson.M{"$set": bson.M{"name": account.Name, "folder": account.Folder, "interval": account.Interval}})
	return err
}

func (r *AccountsRepo) Delete(ctx context.Context, accountID primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": accountID})
	return err
}

func (r *AccountsRepo) GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error {
	var accounts []domain.Account

	cur, err := r.db.Find(ctx, bson.M{"folder": folderID})
	if err != nil {
		return err
	}

	if err := cur.All(ctx, &accounts); err != nil {
		return err
	}

	for _, account := range accounts {
		var updateAccount domain.AccountUpdate
		updateAccount.ID = account.ID
		updateAccount.Name = account.Name
		updateAccount.Folder = account.Folder
		updateAccount.Interval = RandomInterval()

		if err := r.Update(ctx, updateAccount); err != nil {
			return err
		}
	}

	return nil
}

// func (r *AccountsRepo) AddRandomHash(ctx context.Context, accountID primitive.ObjectID, randomHash string) error {
// 	_, err := r.db.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": bson.M{"random_hash": randomHash}})
// 	return err
// }

// func (r *AccountsRepo) AddPhoneHash(ctx context.Context, accountID primitive.ObjectID, phoneCodeHash string) error {
// 	_, err := r.db.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": bson.M{"phone_code_hash": phoneCodeHash}})
// 	return err
// }

// func (r *AccountsRepo) AddApi(ctx context.Context, accountSettings domain.AccountApi) error {
// 	_, err := r.db.UpdateOne(ctx, bson.M{"_id": accountSettings.ID}, bson.M{"$set": bson.M{"api_id": accountSettings.ApiId, "api_hash": accountSettings.ApiHash}})
// 	return err
// }

func (r *AccountsRepo) ChangeStatusBlock(ctx context.Context, accountID primitive.ObjectID, status string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": bson.M{"status_block": status}})
	return err
}

// func (r *AccountsRepo) ChangeVerify(ctx context.Context, accountID primitive.ObjectID) error {
// 	_, err := r.db.UpdateOne(ctx, bson.M{"_id": accountID}, bson.M{"$set": bson.M{"verify": true}})
// 	return err
// }

func (r *AccountsRepo) GetGroupById(ctx context.Context, folderID primitive.ObjectID) (string, error) {
	var folder domain.Folder

	err := r.db.Database().Collection(foldersCollection).FindOne(ctx, bson.M{"_id": folderID}).Decode(&folder)
	if err != nil {
		return "", err
	}

	return folder.Chat, nil
}
