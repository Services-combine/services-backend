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
	"github.com/korpgoodness/service.git/pkg/logging"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	path_python = "/usr/bin/python3"
)

type AccountsService struct {
	repo repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	logger := logging.GetLogger()
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Error loading env variables: %s", err.Error())
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

func (s *AccountsService) CheckingUniqueness(ctx context.Context, phone string) (bool, error) {
	status, err := s.repo.CheckingUniqueness(ctx, phone)
	return status, err
}

func (s *AccountsService) Create(ctx context.Context, accountCreate domain.Account) error {
	accountCreate.Interval = RandomInterval()
	accountCreate.Launch = false
	accountCreate.Status_block = "clean"

	err := s.repo.Create(ctx, accountCreate)
	return err
}

func (s *AccountsService) Update(ctx context.Context, account domain.AccountUpdate) error {
	err := s.repo.Update(ctx, account)
	return err
}

func (s *AccountsService) Delete(ctx context.Context, accountID primitive.ObjectID) error {
	account, err := s.repo.GetById(ctx, accountID)
	if err != nil {
		return err
	}

	if err = s.repo.Delete(ctx, accountID); err != nil {
		return err
	}

	os.Remove(os.Getenv("FOLDER_ACCOUNTS") + account.Phone + ".session")
	return nil
}

func (s *AccountsService) GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repo.GenerateInterval(ctx, folderID)
	return err
}

func (s *AccountsService) CheckBlock(ctx context.Context, folderID primitive.ObjectID) error {
	accounts, err := s.repo.GetAccountsByFolderID(ctx, folderID)
	if err != nil {
		return err
	}

	fmt.Println("CheckBlock")

	for _, account := range accounts {
		fmt.Println(account)
		if account.Api_id != 0 && account.Api_hash != "" {
			script := os.Getenv("FOLDER_PYTHON_SCRIPTS") + "check_block.py"
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

func (s *AccountsService) JoinGroup(ctx context.Context, folderID primitive.ObjectID) error {
	accounts, err := s.repo.GetAccountsByFolderID(ctx, folderID)
	if err != nil {
		return err
	}

	group, err := s.repo.GetGroupById(ctx, folderID)
	if err != nil {
		return err
	}

	for _, account := range accounts {
		if account.Api_id != 0 && account.Api_hash != "" {
			script := os.Getenv("FOLDER_PYTHON_SCRIPTS") + "join_channel.py"
			args_phone := fmt.Sprintf("-P %s", account.Phone)
			args_hash := fmt.Sprintf("-H %s", account.Api_hash)
			args_id := fmt.Sprintf("-I %d", account.Api_id)
			args_group := fmt.Sprintf("-G %s", group)

			status, err := exec.Command(path_python, script, args_phone, args_hash, args_id, args_group).Output()
			if err != nil {
				return err
			}
			fmt.Println(string(status))
		}
	}

	return nil
}
