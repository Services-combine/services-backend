package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountsService struct {
	repo repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	return &AccountsService{repo: repo}
}

func RandomInterval() uint8 {
	min := 15
	max := 40
	rand.Seed(time.Now().UnixNano())
	interval := min + rand.Intn(max-min+1)

	return uint8(interval)
}

func (s *AccountsService) Create(ctx context.Context, accountCreate domain.Account) error {
	accountCreate.Interval = RandomInterval()
	accountCreate.Verify = false
	accountCreate.Launch = false
	accountCreate.Status_block = "clean"

	err := s.repo.Create(ctx, accountCreate)
	return err
}

func (s *AccountsService) GetSettings(ctx context.Context, folderID, accountID primitive.ObjectID) (domain.AccountSettings, error) {
	account, err := s.repo.GetSettings(ctx, accountID)
	if err != nil {
		return domain.AccountSettings{}, err
	}

	folder, err := s.repo.GetFolderByID(ctx, folderID)
	if err != nil {
		return domain.AccountSettings{}, err
	}
	account.FolderName = folder.Name
	account.FolderID = folderID.Hex()
	account.Chat = folder.Chat

	foldersMove := []domain.DataFolderHash{}
	folders, err := s.repo.GetFolders(ctx)
	if err != nil {
		return domain.AccountSettings{}, err
	}

	for Name, ObjectID := range folders {
		if ObjectID != folderID.Hex() {
			foldersMove = append(foldersMove, domain.DataFolderHash{Name, ObjectID})
		}
	}
	account.FoldersMove = foldersMove

	return account, nil
}

func (s *AccountsService) UpdateAccount(ctx context.Context, account domain.AccountUpdate) error {
	err := s.repo.UpdateAccount(ctx, account)
	return err
}

func (s *AccountsService) Delete(ctx context.Context, accountID primitive.ObjectID) error {
	err := s.repo.Delete(ctx, accountID)
	return err
}

func (s *AccountsService) GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repo.GenerateInterval(ctx, folderID)
	return err
}

func (s *AccountsService) CheckBlock(ctx context.Context, folderID primitive.ObjectID) error {
	accounts, err := s.repo.GetAccountsFolder(ctx, folderID)
	if err != nil {
		return err
	}

	for _, account := range accounts {
		if account.Api_id != 0 && account.Api_hash != "" && account.Verify {
			fmt.Println(account)
		}
	}

	return nil
}

func GetStatusBlock(ctx context.Context, account domain.Account) (string, error) {
	client := telegram.NewClient(account.Api_id, account.Api_hash, telegram.Options{})

	if err := client.Run(ctx, func(ctx context.Context) error {
		api := client.API()
		fmt.Println(api)

		return nil
	}); err != nil {
		return "", err
	}

	return "", nil
}
