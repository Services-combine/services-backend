package service

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	path_accounts = "/home/q/p/projects/services/backend/accounts/"
)

type AccountsService struct {
	repo repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

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
	account, err := s.repo.GetData(ctx, accountID)
	if err != nil {
		return err
	}

	if err = s.repo.Delete(ctx, accountID); err != nil {
		return err
	}

	os.Remove(path_accounts + account.Phone + ".session")
	return nil
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
			script := os.Getenv("FOLDER_PYTHON_SCRIPTS_VERIFY") + "check_block.py"
			args_phone := fmt.Sprintf("-P %s", account.Phone)
			args_hash := fmt.Sprintf("-H %s", account.Api_hash)
			args_id := fmt.Sprintf("-I %d", account.Api_id)

			status, err := exec.Command(path_python, script, args_phone, args_hash, args_id).Output()
			if err != nil {
				return err
			}

			if string(status) != "ERROR" {
				err = s.repo.ChangeStatusBlock(ctx, account.ID, string(status))
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
