package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" binding:"required"`
	Phone        string             `json:"phone" binding:"required"`
	Folder       primitive.ObjectID `json:"folder"`
	Api_id       int                `json:"api_id"`
	Api_hash     string             `json:"api_hash"`
	Verify       bool               `json:"verify"`
	Launch       bool               `json:"launch"`
	Interval     uint8              `json:"interval"`
	Status_block string             `json:"status_block"`
}

type AccountSettings struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" binding:"required"`
	Phone        string             `json:"phone" binding:"required"`
	Launch       bool               `json:"launch"`
	Interval     uint8              `json:"interval"`
	Status_block string             `json:"status_block"`
	Folder_name  string             `json:"folder_name"`
	Chat         string             `json:"chat"`
}
