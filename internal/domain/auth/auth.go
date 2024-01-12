package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Session struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	UserAgent    string             `json:"user_agent" bson:"user_agent"`
	ClientIP     string             `json:"client_ip" bson:"client_ip"`
	IsBlocked    bool               `json:"is_blocked" bson:"is_blocked"`
	ExpiresAt    time.Time          `json:"expires_at" bson:"expires_at"`
}
