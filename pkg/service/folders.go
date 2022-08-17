package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MODE_INVITING          = "inviting"
	MODE_MAILING_USERNAMES = "mailing-usernames"
	MODE_MAILING_GROUPS    = "mailing-groups"
)

type FoldersService struct {
	repo repository.Folders
}

func NewFoldersService(repo repository.Folders) *FoldersService {
	return &FoldersService{repo: repo}
}

func GenerateHash() string {
	const LENGTH_HASH = 34
	const symbols = "1234567890qwertyuiopasdfghjklzxcvbnm"
	random_hash := make([]byte, LENGTH_HASH)

	rand.Seed(time.Now().UnixNano())
	for i := range random_hash {
		random_hash[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(random_hash)
}

func (s *FoldersService) GetDataMainPage(ctx context.Context) (map[string]interface{}, error) {
	dataPage := map[string]interface{}{}

	folders, err := s.GetListFolders(ctx, "/")
	if err != nil {
		return nil, err
	}
	dataPage["folders"] = folders

	countAccounts, err := s.repo.GetCountAllAccount(ctx)
	if err != nil {
		return nil, err
	}
	dataPage["countAccounts"] = countAccounts

	return dataPage, nil
}

func (s *FoldersService) GetListFolders(ctx context.Context, path string) ([]domain.FolderItem, error) {
	folders, err := s.repo.GetListFolders(ctx, path)
	return folders, err
}

func (s *FoldersService) Get(ctx context.Context, path string) ([]domain.Folder, error) {
	folders, err := s.repo.Get(ctx, path)
	return folders, err
}

func (s *FoldersService) Create(ctx context.Context, folder domain.Folder) error {
	err := s.repo.Create(ctx, folder)
	return err
}

func (s *FoldersService) GetData(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) {
	folder, err := s.repo.GetData(ctx, folderID)
	if err != nil {
		return domain.Folder{}, err
	}

	return folder, nil
}

func (s *FoldersService) GetFolderById(ctx context.Context, folderID primitive.ObjectID) (map[string]interface{}, error) {
	folderData := map[string]interface{}{}

	folder, err := s.GetData(ctx, folderID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["folder"] = folder

	accounts, err := s.repo.GetAccountsByFolderID(ctx, folderID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["accounts"] = accounts

	accountsMove := []domain.DataFolderHash{}
	foldersAll, err := s.repo.GetFolders(ctx)
	if err != nil {
		return map[string]interface{}{}, err
	}
	for _, folder := range foldersAll {
		if folder.ID.Hex() != folderID.Hex() {
			accountsMove = append(accountsMove, domain.DataFolderHash{folder.Name, folder.ID.Hex()})
		}
	}
	folderData["accountsMove"] = accountsMove

	folders, err := s.GetListFolders(ctx, folderID.Hex())
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["folders"] = folders

	countAccounts, err := s.repo.GetCountAccounts(ctx, folderID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["countAccounts"] = countAccounts

	pathHash, err := GetPathHash(ctx, folderID, folder.Path, s.repo)
	if err != nil {
		return nil, err
	}
	folderData["pathHash"] = pathHash

	return folderData, nil
}

func ConvertPath(path string) (primitive.ObjectID, error) {
	ObjectID, err := primitive.ObjectIDFromHex(path)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return ObjectID, nil
}

func (s *FoldersService) GetFoldersMove(ctx context.Context, folderID primitive.ObjectID) ([]domain.DataFolderHash, error) {
	folder, err := s.GetData(ctx, folderID)
	if err != nil {
		return nil, err
	}
	path := folder.Path

	foldersMove := map[string]string{}
	status := 0

	if path != "/" {
		ObjectID, err := ConvertPath(path)
		if err != nil {
			return nil, err
		}
		mainFolder, err := s.repo.GetData(ctx, ObjectID)
		if err != nil {
			return nil, err
		}
		foldersMove[mainFolder.Name] = path
	} else {
		foldersMove["/"] = "/"
	}

	folders, err := s.repo.GetFolders(ctx)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		if folderID != folder.ID && path != folder.ID.Hex() {
			nextPath := folder.Path
			nextFolderID := folder.ID
			status = 0

			for nextPath != "/" {
				if nextPath == folderID.Hex() {
					status = 1
					break
				}

				nextPathObject, err := ConvertPath(nextPath)
				if err != nil {
					return nil, err
				}
				nextFolder, err := s.repo.GetData(ctx, nextPathObject)
				if err != nil {
					return nil, err
				}
				nextFolderID = nextFolder.ID

				nextFolder, err = s.repo.GetData(ctx, nextFolderID)
				if err != nil {
					return nil, err
				}
				nextPath = nextFolder.Path
			}

			if status == 0 {
				foldersMove[folder.Name] = folder.ID.Hex()
			}
		}
	}

	if _, found := foldersMove["/"]; !found {
		foldersMove["/"] = "/"
	}

	MapFoldersMove := []domain.DataFolderHash{}
	for Name, ObjectID := range foldersMove {
		if path != ObjectID {
			MapFoldersMove = append(MapFoldersMove, domain.DataFolderHash{Name, ObjectID})
		}
	}

	return MapFoldersMove, nil
}

func GetPathHash(ctx context.Context, folderID primitive.ObjectID, path string, db repository.Folders) ([]domain.DataFolderHash, error) {
	foldersHash := map[string]string{}
	MapFoldersHash := []domain.DataFolderHash{}

	folders, err := db.GetFolders(ctx)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		foldersHash[folder.Name] = folder.ID.Hex()
	}

	for {
		nextFolder, err := db.GetData(ctx, folderID)
		if err != nil {
			return nil, err
		}

		MapFoldersHash = append(MapFoldersHash, domain.DataFolderHash{nextFolder.Name, nextFolder.ID.Hex()})
		if nextFolder.Path == "/" {
			break
		}

		folderID, err = ConvertPath(nextFolder.Path)
		if err != nil {
			return nil, err
		}
	}

	ReverceFoldersHash := ReverseSlice(MapFoldersHash)
	return ReverceFoldersHash, nil
}

func ReverseSlice(s []domain.DataFolderHash) []domain.DataFolderHash {
	a := make([]domain.DataFolderHash, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func (s *FoldersService) Move(ctx context.Context, folderID primitive.ObjectID, path string) error {
	err := s.repo.Move(ctx, folderID, path)
	return err
}

func (s *FoldersService) Rename(ctx context.Context, folderID primitive.ObjectID, name string) error {
	err := s.repo.Rename(ctx, folderID, name)
	return err
}

func (s *FoldersService) ChangeChat(ctx context.Context, folderID primitive.ObjectID, chat string) error {
	err := s.repo.ChangeChat(ctx, folderID, chat)
	return err
}

func (s *FoldersService) ChangeUsernames(ctx context.Context, folderID primitive.ObjectID, usernames []string) error {
	err := s.repo.ChangeUsernames(ctx, folderID, usernames)
	return err
}

func (s *FoldersService) ChangeMessage(ctx context.Context, folderID primitive.ObjectID, message string) error {
	err := s.repo.ChangeMessage(ctx, folderID, message)
	return err
}

func (s *FoldersService) ChangeGroups(ctx context.Context, folderID primitive.ObjectID, groups []string) error {
	err := s.repo.ChangeGroups(ctx, folderID, groups)
	return err
}

func (s *FoldersService) Delete(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repo.Delete(ctx, folderID)
	return err
}

func (s *FoldersService) CheckingEnteredData(ctx context.Context, folderID primitive.ObjectID, mode string) error {
	folderData, err := s.repo.GetData(ctx, folderID)
	if err != nil {
		return err
	}

	settings, err := s.repo.GetSettings(ctx)
	if err != nil {
		return err
	}

	accounts, err := s.repo.GetAccountsByFolderID(ctx, folderID)
	if err != nil {
		return err
	}

	checkInternal := 0
	for _, account := range accounts {
		if account.Interval != 0 {
			checkInternal++
		}
	}

	if len(folderData.Usernames) == 0 {
		if mode == MODE_INVITING || mode == MODE_MAILING_USERNAMES {
			return fmt.Errorf("First specify the usernames")
		}
	} else if folderData.Chat == "" {
		if mode == MODE_INVITING {
			return fmt.Errorf("First specify the chat")
		}
	} else if len(folderData.Usernames) < (len(accounts) * settings.CountInviting) {
		if mode == MODE_INVITING {
			return fmt.Errorf("The number of usernames is not enough for all accounts")
		}
	} else if len(folderData.Usernames) < (len(accounts) * settings.CountMailing) {
		if mode == MODE_MAILING_USERNAMES {
			return fmt.Errorf("The number of usernames is not enough for all accounts")
		}
	} else if folderData.Message == "" {
		if mode == MODE_MAILING_GROUPS || mode == MODE_MAILING_USERNAMES {
			return fmt.Errorf("First specify the message")
		}
	} else if len(folderData.Groups) == 0 {
		if mode == MODE_MAILING_GROUPS {
			return fmt.Errorf("First specify the groups")
		}
	} else if checkInternal == 0 {
		return fmt.Errorf("The %d accounts do not have intervals set", checkInternal)
	}

	return nil
}

func (s *FoldersService) LaunchInviting(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.CheckingEnteredData(ctx, folderID, MODE_INVITING)
	if err != nil {
		return err
	}

	err = s.repo.LaunchInviting(ctx, folderID)
	return err
}

func (s *FoldersService) LaunchMailingUsernames(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.CheckingEnteredData(ctx, folderID, MODE_MAILING_USERNAMES)
	if err != nil {
		return err
	}

	err = s.repo.LaunchMailingUsernames(ctx, folderID)
	return err
}

func (s *FoldersService) LaunchMailingGroups(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.CheckingEnteredData(ctx, folderID, MODE_MAILING_GROUPS)
	if err != nil {
		return err
	}

	err = s.repo.LaunchMailingGroups(ctx, folderID)
	return err
}
