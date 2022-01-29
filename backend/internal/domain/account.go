package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" binding:"required"`
	Phone         string             `json:"phone" binding:"required"`
	Folder        primitive.ObjectID `json:"folder"`
	Api_id        int                `json:"api_id"`
	Api_hash      string             `json:"api_hash"`
	Verify        bool               `json:"verify"`
	Launch        bool               `json:"launch"`
	Interval      uint8              `json:"interval"`
	Status_block  string             `json:"status_block"`
	Random_hash   string             `json:"random_hash"`
	PhoneCodeHash string             `json:"phone_code_hash"`
}

type AccountSettings struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" binding:"required"`
	Phone        string             `json:"phone" binding:"required"`
	Launch       bool               `json:"launch"`
	Interval     uint8              `json:"interval"`
	Status_block string             `json:"status_block"`
	FolderName   string             `json:"folder_name"`
	FolderID     string             `json:"folder_id"`
	FoldersMove  map[string]string  `json:"folders_move"`
	Chat         string             `json:"chat"`
}

type AccountUpdate struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" binding:"required"`
	FolderID string             `json:"folder_id" binding:"required"`
	Folder   primitive.ObjectID `json:"folder"`
	Interval uint8              `json:"interval" binding:"required"`
}

type AccountsCount struct {
	All   int `json:"all"`
	Clean int `json:"clean"`
	Block int `json:"block"`
}

type AccountLogin struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Password string             `json:"password"`
}

type AccountApi struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	ApiId   int                `json:"api_id"`
	ApiHash string             `json:"api_hash"`
}
