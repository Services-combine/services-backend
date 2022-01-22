package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *FoldersService) Get(ctx context.Context, path string) ([]domain.Folder, error) {
	folders, err := s.repo.Get(ctx, path)
	return folders, err
}

func (s *FoldersService) Create(ctx context.Context, folder domain.Folder) error {
	err := s.repo.Create(ctx, folder)
	return err
}

func (s *FoldersService) OpenFolder(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, []domain.Account, domain.AccountsCount, map[string]string, error) {
	var accounts []domain.Account
	var countAccounts domain.AccountsCount

	folder, err := s.repo.GetData(ctx, folderID)
	if err != nil {
		return domain.Folder{}, nil, domain.AccountsCount{}, nil, err
	}

	accounts, err = s.repo.GetAccountByFolderID(ctx, folderID)
	if err != nil {
		return folder, nil, domain.AccountsCount{}, nil, err
	}

	countAccounts, err = s.repo.GetCountAccounts(ctx, folderID)
	if err != nil {
		return folder, accounts, domain.AccountsCount{}, nil, err
	}

	foldersMove, err := GetFoldersMove(ctx, folderID, folder.Path, s.repo)
	if err != nil {
		return folder, accounts, countAccounts, nil, err
	}

	return folder, accounts, countAccounts, foldersMove, nil
}

func ConvertPath(path string) (primitive.ObjectID, error) {
	ObjectID, err := primitive.ObjectIDFromHex(path)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return ObjectID, nil
}

func GetFoldersMove(ctx context.Context, folderID primitive.ObjectID, path string, db repository.Folders) (map[string]string, error) {
	foldersMove := map[string]string{}

	if path != "/" {
		ObjectID, err := ConvertPath(path)
		if err != nil {
			return nil, err
		}
		mainFolder, err := db.GetData(ctx, ObjectID)
		if err != nil {
			return nil, err
		}
		foldersMove[mainFolder.Name] = path
	} else {
		foldersMove["/"] = "/"
	}

	folders, err := db.GetFolders(ctx)
	if err != nil {
		return nil, err
	}

	for _, folder := range folders {
		if folderID != folder.ID && path != folder.ID.Hex() {
			path_ := folder.Path
			folderID_ := folder.ID
			status := 0

			for {
				switch path_ {
				case "/":
					break
				case folderID_.Hex():
					status = 1
					break
				}

				nextFolder, err := db.GetData(ctx, folderID_)
				if err != nil {
					return nil, err
				}
				path_ = nextFolder.Path

				pathObject, err := ConvertPath(path_)
				if err != nil {
					return nil, err
				}
				nextFolder, err = db.GetData(ctx, pathObject)
				if err != nil {
					return nil, err
				}
				folderID_ = nextFolder.ID
			}

			if status == 0 {
				foldersMove[folder.Name] = folder.ID.Hex()
			}
		}
	}

	if _, found := foldersMove["/"]; !found {
		foldersMove["/"] = "/"
	}

	return foldersMove, nil
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

func (s *FoldersService) LaunchInviting(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repo.LaunchInviting(ctx, folderID)
	return err
}

func (s *FoldersService) LaunchMailingUsernames(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repo.LaunchMailingUsernames(ctx, folderID)
	return err
}

func (s *FoldersService) LaunchMailingGroups(ctx context.Context, folderID primitive.ObjectID) error {
	err := s.repo.LaunchMailingGroups(ctx, folderID)
	return err
}
