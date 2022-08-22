package domain

type ChannelAdd struct {
	ChannelId       string `json:"channel_id" binding:"required"`
	ApiKey          string `json:"api_key" binding:"required"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Photo           string `json:"photo"`
	VideoCount      uint64 `json:"video_count"`
	ViewCount       uint64 `json:"view_count"`
	SubscriberCount uint64 `json:"subscriber_count"`
}
