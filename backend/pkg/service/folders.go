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

func (s *FoldersService) OpenFolder(ctx context.Context, folderID primitive.ObjectID) (map[string]interface{}, error) {
	folderData := map[string]interface{}{}

	folder, err := s.repo.GetData(ctx, folderID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["folder"] = folder

	accounts, err := s.repo.GetAccountByFolderID(ctx, folderID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["accounts"] = accounts

	countAccounts, err := s.repo.GetCountAccounts(ctx, folderID)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["countAccounts"] = countAccounts

	foldersMove, err := GetFoldersMove(ctx, folderID, folder.Path, s.repo)
	if err != nil {
		return map[string]interface{}{}, err
	}
	folderData["foldersMove"] = foldersMove

	pathHash, err := GetpathHash(ctx, folderID, folder.Path, s.repo)
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

func GetFoldersMove(ctx context.Context, folderID primitive.ObjectID, path string, db repository.Folders) (map[string]string, error) {
	foldersMove := map[string]string{}
	status := 0

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
				nextFolder, err := db.GetData(ctx, nextPathObject)
				if err != nil {
					return nil, err
				}
				nextFolderID = nextFolder.ID

				nextFolder, err = db.GetData(ctx, nextFolderID)
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

	return foldersMove, nil
}

func GetpathHash(ctx context.Context, folderID primitive.ObjectID, path string, db repository.Folders) (map[string]string, error) {
	foldersHash := map[string]string{}
	pathHash := map[string]string{}

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

		if nextFolder.Path == "/" {
			pathHash[nextFolder.Name] = nextFolder.ID.Hex()
			break
		}
		pathHash[nextFolder.Name] = nextFolder.ID.Hex()
		folderID, err = ConvertPath(nextFolder.Path)
		if err != nil {
			return nil, err
		}
	}

	return pathHash, nil
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
