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
	Random_hash  string             `json:"random_hash"`
}

type AccountSettings struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" binding:"required"`
	Phone        string             `json:"phone" binding:"required"`
	Launch       bool               `json:"launch"`
	Interval     uint8              `json:"interval"`
	Status_block string             `json:"status_block"`
	Folder_name  string             `json:"folder_name"`
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
	CountAll   int `json:"count_all"`
	CountClean int `json:"count_clean"`
	CountBlock int `json:"count_block"`
}

type AccountLogin struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Phone       string             `json:"phone"`
	Random_hash string             `json:"random_hash"`
	Password    string             `json:"password"`
}

type AccountUserAgent struct {
	UserAgent string `json:"user_agent"`
}

type AccountCreateApp struct {
	Hash         string `json:"hash"`
	AppTitle     string `json:"app_title"`
	AppShortname string `json:"app_shortname"`
	AppUrl       string `json:"app_url"`
	AppPlatform  string `json:"app_platform"`
	AppDesc      string `json:"app_desc"`
}
