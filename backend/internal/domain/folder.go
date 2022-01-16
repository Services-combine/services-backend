package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Folder struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name              string             `json:"name" binding:"required"`
	Hash              string             `json:"hash"`
	Path              string             `json:"path"`
	Chat              string             `json:"chat"`
	Usernames         []string           `json:"usernames"`
	Message           string             `json:"message"`
	Groups            []string           `json:"groups"`
	Inviting          bool               `json:"inviting"`
	Mailing_usernames bool               `json:"mailing_usernames"`
	Mailing_groups    bool               `json:"mailing_groups"`
}

type FolderMove struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

type FolderRename struct {
	Name string `json:"name" binding:"required"`
}

type FolderChat struct {
	Chat string `json:"chat" binding:"required"`
}

type FolderUsernames struct {
	Usernames []string `json:"usernames" binding:"required"`
}

type FolderMessage struct {
	Message string `json:"message" binding:"required"`
}

type FolderGroups struct {
	Groups []string `json:"groups" binding:"required"`
}
