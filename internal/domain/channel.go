package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChannelIdKey struct {
	ChannelId string `json:"channel_id" binding:"required"`
	ApiKey    string `json:"api_key"`
}

type ChannelAdd struct {
	ChannelId       string `json:"channel_id" binding:"required"`
	ApiKey          string `json:"api_key" binding:"required"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Photo           string `json:"photo"`
	VideoCount      uint64 `json:"video_count"`
	ViewCount       uint64 `json:"view_count"`
	SubscriberCount uint64 `json:"subscriber_count"`
	Launch          bool   `json:"launch"`
}

type ChannelGet struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChannelId            string             `json:"channel_id" binding:"required"`
	ApiKey               string             `json:"api_key" binding:"required"`
	Title                string             `json:"title"`
	Description          string             `json:"description"`
	Photo                string             `json:"photo"`
	VideoCount           uint64             `json:"video_count"`
	ViewCount            uint64             `json:"view_count"`
	SubscriberCount      uint64             `json:"subscriber_count"`
	Launch               bool               `json:"launch"`
	Comment              string             `json:"comment"`
	CountCommentedVideos uint32             `json:"count_commented_videos"`
}

type ChannelEdit struct {
	Comment              string `json:"comment"`
	CountCommentedVideos uint32 `json:"count_commented_videos"`
}
