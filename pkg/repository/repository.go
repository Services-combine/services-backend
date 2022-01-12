package repository

import "go.mongodb.org/mongo-driver/mongo"

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{}
}
