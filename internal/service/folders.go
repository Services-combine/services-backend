package service

import (
	"context"
	"fmt"
	"github.com/b0shka/services/internal/domain"
	domain_folders "github.com/b0shka/services/internal/domain/folders"
	"github.com/b0shka/services/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MODE_INVITING          = "inviting"
	MODE_MAILING_USERNAMES = "mailing-usernames"
	MODE_MAILING_GROUPS    = "mailing-groups"
)

type FoldersService struct {
	repoFolders  repository.Folders
	repoSettings repository.Settings
}

func NewFoldersService(
	repoFolders repository.Folders,
	repoSettings repository.Settings,
) *FoldersService {
	return &FoldersService{
		repoFolders:  repoFolders,
		repoSettings: repoSettings,
	}
}

func (s *FoldersService) GetFolders(ctx context.Context) (domain_folders.GetFoldersOutput, error) {
	folders, err := s.GetFoldersByPath(ctx, "/")
	if err != nil {
		return domain_folders.GetFoldersOutput{}, err
	}

	countAccounts, err := s.repoFolders.GetCountAccounts(ctx, primitive.NilObjectID)
	if err != nil {
		return domain_folders.GetFoldersOutput{}, err
	}

	res := domain_folders.NewGetFoldersOutput(folders, countAccounts)
	return res, nil
}

func (s *FoldersService) Create(ctx context.Context, folder domain.Folder) error {
	err := s.repoFolders.Create(ctx, folder)
	return err
}

func (s *FoldersService) GetFoldersByPath(ctx context.Context, path string) ([]domain.FolderItem, error) {
	folders, err := s.repoFolders.GetFoldersByPath(ctx, path)
	return folders, err
}

func (s *FoldersService) GetFolderById(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) {
	folder, err := s.repoFolders.GetFolderById(ctx, folderID)
	if err != nil {
		return domain.Folder{}, err
	}

	return folder, nil
}

func (s *FoldersService) GetAllDataFolderById(ctx context.Context, folderID primitive.ObjectID) (domain_folders.GetFolderOutput, error) {
	folder, err := s.GetFolderById(ctx, folderID)
	if err != nil {
		return domain_folders.GetFolderOutput{}, err
	}

	accounts, err := s.repoFolders.GetAccountsByFolderID(ctx, folderID)
	if err != nil {
		return domain_folders.GetFolderOutput{}, err
	}

	accountsMove := []domain.AccountDataMove{}
	foldersAll, err := s.repoFolders.GetFolders(ctx)
	if err != nil {
		return domain_folders.GetFolderOutput{}, err
	}
	for _, folder := range foldersAll {
		if folder.ID.Hex() != folderID.Hex() {
			accountsMove = append(accountsMove, domain.AccountDataMove{folder.Name, folder.ID.Hex()})
		}
	}

	folders, err := s.GetFoldersByPath(ctx, folderID.Hex())
	if err != nil {
		return domain_folders.GetFolderOutput{}, err
	}

	countAccounts, err := s.repoFolders.GetCountAccounts(ctx, folderID)
	if err != nil {
		return domain_folders.GetFolderOutput{}, err
	}

	pathHash, err := GetPathHash(ctx, folderID, folder.Path, s.repoFolders)
	if err != nil {
		return domain_folders.GetFolderOutput{}, err
	}

	folderData := domain_folders.NewGetFolderOutput(folder, accounts, accountsMove, folders, countAccounts, pathHash)
	return folderData, nil
}

func ConvertPath(path string) (primitive.ObjectID, error) {
	ObjectID, err := primitive.ObjectIDFromHex(path)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return ObjectID, nil
}

func (s *FoldersService) GetFoldersMove(ctx context.Context, folderID primitive.ObjectID) ([]domain.AccountDataMove, error) {
	folder, err := s.GetFolderById(ctx, folderID)
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
		mainFolder, err := s.repoFolders.GetFolderById(ctx, ObjectID)
		if err != nil {
			return nil, err
		}
		foldersMove[mainFolder.Name] = path
	} else {
		foldersMove["/"] = "/"
	}

	folders, err := s.repoFolders.GetFolders(ctx)
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
				nextFolder, err := s.repoFolders.GetFolderById(ctx, nextPathObject)
				if err != nil {
					return nil, err
				}
				nextFolderID = nextFolder.ID

				nextFolder, err = s.repoFolders.GetFolderById(ctx, nextFolderID)
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

	MapFoldersMove := []domain.AccountDataMove{}
	for Name, ObjectID := range foldersMove {
		if path != ObjectID {
			MapFoldersMove = append(MapFoldersMove, domain.AccountDataMove{Name, ObjectID})
		}
	}

	return MapFoldersMove, nil
}

func GetPathHash(ctx context.Context, folderID primitive.ObjectID, path string, db repository.Folders) ([]domain.AccountDataMove, error) {
	foldersHash := map[string]string{}
	MapFoldersHash := []domain.AccountDataMove{}

	folders, err := db.GetFolders(ctx)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		foldersHash[folder.Name] = folder.ID.Hex()
	}

	for {
		nextFolder, err := db.GetFolderById(ctx, folderID)
		if err != nil {
			return nil, err
		}

		MapFoldersHash = append(MapFoldersHash, domain.AccountDataMove{nextFolder.Name, nextFolder.ID.Hex()})
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

func ReverseSlice(s []domain.AccountDataMove) []domain.AccountDataMove {
	a := make([]domain.AccountDataMove, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func (s *FoldersService) Move(ctx context.Context, folderID primitive.ObjectID, path string) error {
	err := s.repoFolders.Move(ctx, folderID, path)
	return err
}

func (s *FoldersService) Rename(ctx context.Context, folderID primitive.ObjectID, name string) error {
	err := s.repoFolders.Rename(ctx, folderID, name)
	return err
}

func (s *FoldersService) ChangeChat(ctx context.Context, folderID primitive.ObjectID, chat string) error {
	err := s.repoFolders.ChangeChat(ctx, folderID, chat)
	return err
}

func (s *FoldersService) ChangeUsernames(ctx context.Context, folderID primitive.ObjectID, usernames []string) error {
	err := s.repoFolders.ChangeUsernames(ctx, folderID, usernames)
	return err
}

func (s *FoldersService) ChangeMessage(ctx context.Context, folderID primitive.ObjectID, message string) error {
	err := s.repoFolders.ChangeMessage(ctx, folderID, message)
	return err
}

func (s *FoldersService) ChangeGroups(ctx context.Context, folderID primitive.ObjectID, groups []string) error {
	err := s.repoFolders.ChangeGroups(ctx, folderID, groups)
	return err
}

func (s *FoldersService) Delete(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repoFolders.Delete(ctx, folderID)
	return err
}

func (s *FoldersService) CheckingEnteredData(ctx context.Context, folderID primitive.ObjectID, mode string) error {
	folderData, err := s.repoFolders.GetFolderById(ctx, folderID)
	if err != nil {
		return err
	}

	settings, err := s.repoSettings.Get(ctx, repository.ServiceInviting)
	if err != nil {
		return err
	}

	accounts, err := s.repoFolders.GetAccountsByFolderID(ctx, folderID)
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

	err = s.repoFolders.LaunchInviting(ctx, folderID)
	return err
}

func (s *FoldersService) LaunchMailingUsernames(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.CheckingEnteredData(ctx, folderID, MODE_MAILING_USERNAMES)
	if err != nil {
		return err
	}

	err = s.repoFolders.LaunchMailingUsernames(ctx, folderID)
	return err
}

func (s *FoldersService) LaunchMailingGroups(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.CheckingEnteredData(ctx, folderID, MODE_MAILING_GROUPS)
	if err != nil {
		return err
	}

	err = s.repoFolders.LaunchMailingGroups(ctx, folderID)
	return err
}
