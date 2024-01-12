package folders

import "github.com/b0shka/services/internal/domain"

type GetFoldersOutput struct {
	Folders       []domain.FolderItem  `json:"folders"`
	CountAccounts domain.AccountsCount `json:"count_accounts"`
}

func NewGetFoldersOutput(
	folders []domain.FolderItem,
	countAccounts domain.AccountsCount,
) GetFoldersOutput {
	return GetFoldersOutput{
		folders,
		countAccounts,
	}
}

type GetFolderOutput struct {
	Folder        domain.Folder            `json:"folder"`
	Accounts      []domain.Account         `json:"accounts"`
	AccountsMove  []domain.AccountDataMove `json:"accounts_move"`
	Folders       []domain.FolderItem      `json:"folders"`
	CountAccounts domain.AccountsCount     `json:"count_accounts"`
	PathHash      []domain.AccountDataMove `json:"path_hash"`
}

func NewGetFolderOutput(
	folder domain.Folder,
	accounts []domain.Account,
	accountsMove []domain.AccountDataMove,
	folders []domain.FolderItem,
	countAccounts domain.AccountsCount,
	pathHash []domain.AccountDataMove,
) GetFolderOutput {
	return GetFolderOutput{
		folder,
		accounts,
		accountsMove,
		folders,
		countAccounts,
		pathHash,
	}
}
