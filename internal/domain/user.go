package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
}
