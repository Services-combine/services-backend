package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Folder struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id"`
	Name              string             `json:"name" binding:"required"`
	Hash              string             `json:"hash"`
	Path              string             `json:"path" binding:"required"`
	Chat              string             `json:"chat"`
	Usernames         []string           `json:"usernames"`
	Messages          []string           `json:"messages"`
	Message_command   string             `json:"message_command"`
	Groups            []string           `json:"groups"`
	Inviting          bool               `json:"inviting"`
	Mailing_usernames bool               `json:"mailing_usernames"`
	Mailing_groups    bool               `json:"mailing_groups"`
}
