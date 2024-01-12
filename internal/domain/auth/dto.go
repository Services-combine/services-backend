package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginInput(email, password string) LoginInput {
	return LoginInput{
		Email:    email,
		Password: password,
	}
}

type LoginOutput struct {
	SessionID    primitive.ObjectID `json:"session_id"`
	RefreshToken string             `json:"refresh_token"`
	AccessToken  string             `json:"access_token"`
}

func NewLoginOutput(sessionID primitive.ObjectID, refreshToken, accessToken string) LoginOutput {
	return LoginOutput{
		SessionID:    sessionID,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token"`
}

func NewRefreshTokenInput(refreshToken string) RefreshTokenInput {
	return RefreshTokenInput{
		RefreshToken: refreshToken,
	}
}

type RefreshTokenOutput struct {
	AccessToken string `json:"access_token"`
}

func NewRefreshTokenOutput(accessToken string) RefreshTokenOutput {
	return RefreshTokenOutput{
		AccessToken: accessToken,
	}
}
