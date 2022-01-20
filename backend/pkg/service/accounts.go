package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	link_get_password = "https://my.telegram.org/auth/send_password"
	link_authorized   = "https://my.telegram.org/auth/login"
	link_apps         = "https://my.telegram.org/apps"
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
	var accountSettings domain.AccountSettings
	account, err := s.repo.GetData(ctx, accountID)
	if err != nil {
		return accountSettings, err
	}
	accountSettings.ID = account.ID
	accountSettings.Name = account.Name
	accountSettings.Phone = account.Phone
	accountSettings.Launch = account.Launch
	accountSettings.Interval = account.Interval
	accountSettings.Status_block = account.Status_block

	var folder domain.Folder
	folder, err = s.repo.GetFolderByID(ctx, folderID)
	if err != nil {
		return accountSettings, err
	}
	accountSettings.Folder_name = folder.Name
	accountSettings.Chat = folder.Chat

	return accountSettings, err
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

func (s *AccountsService) LoginApi(ctx context.Context, accountID primitive.ObjectID) error {
	account, err := s.repo.GetData(ctx, accountID)
	if err != nil {
		return err
	}

	data := url.Values{
		"phone": {account.Phone},
	}

	resp, err := http.PostForm(link_get_password, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Login api: %s", resp.Status)
	}

	var getData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&getData)
	if err != nil {
		return err
	}
	randomHash := getData["random_hash"].(string)
	if err := s.repo.AddRandomHash(ctx, accountID, randomHash); err != nil {
		return err
	}

	return nil
}

func (s *AccountsService) ParsingApi(ctx context.Context, accountLogin domain.AccountLogin) error {
	account, err := s.repo.GetData(ctx, accountLogin.ID)
	if err != nil {
		return err
	}

	data := url.Values{
		"phone":       {account.Phone},
		"random_hash": {account.Random_hash},
		"password":    {accountLogin.Password},
	}

	resp, err := http.PostForm(link_authorized, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Login api: %s", resp.Status)
	}

	return nil
}
