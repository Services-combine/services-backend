package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type MarkCreate struct {
	Title string `json:"title" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type MarkGet struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title" binding:"required"`
	Color string             `json:"color" binding:"required"`
}
