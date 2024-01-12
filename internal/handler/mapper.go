package handler

import (
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	domain_folders "github.com/b0shka/services/internal/domain/folders"
	domain_settings "github.com/b0shka/services/internal/domain/settings"
)

func NewLoginInput(req LoginRequest) domain_auth.LoginInput {
	return domain_auth.NewLoginInput(req.Email, req.Password)
}

func NewLoginResponse(out domain_auth.LoginOutput) LoginResponse {
	return LoginResponse{
		SessionID:    out.SessionID,
		RefreshToken: out.RefreshToken,
		AccessToken:  out.AccessToken,
	}
}

func NewRefreshTokenInput(req RefreshTokenRequest) domain_auth.RefreshTokenInput {
	return domain_auth.NewRefreshTokenInput(req.RefreshToken)
}

func NewRefreshTokenResponse(out domain_auth.RefreshTokenOutput) RefreshTokenResponse {
	return RefreshTokenResponse{
		AccessToken: out.AccessToken,
	}
}

func NewSaveSettingsInput(req SaveSettingsRequest) domain_settings.SaveSettingsInput {
	return domain_settings.NewSaveSettingsInput(req.CountInviting, req.CountMailing)
}

func NewGetFoldersResponse(out domain_folders.GetFoldersOutput) GetFoldersResponse {
	return GetFoldersResponse{
		Folders:       out.Folders,
		CountAccounts: out.CountAccounts,
	}
}

func NewGetFolderResponse(out domain_folders.GetFolderOutput) GetFolderResponse {
	return GetFolderResponse{
		Folder:        out.Folder,
		Accounts:      out.Accounts,
		AccountsMove:  out.AccountsMove,
		Folders:       out.Folders,
		CountAccounts: out.CountAccounts,
		PathHash:      out.PathHash,
	}
}
