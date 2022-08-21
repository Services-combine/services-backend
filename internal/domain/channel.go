package domain

type ChannelAdd struct {
	ChannelId       string `json:"channel_id" binding:"required"`
	ApiKey          string `json:"api_key" binding:"required"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Photo           string `json:"photo"`
	VideoCount      int    `json:"video_count"`
	ViewCount       int    `json:"view_count"`
	SubscriberCount int    `json:"subscriber_count"`
}

type ChannelDataApi struct {
	Snippet    SnippetChannel    `json:"snippet" binding:"required"`
	Statistics StatisticsChannel `json:"statistics" binding:"required"`
}

type SnippetChannel struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type StatisticsChannel struct {
	VideoCount      int `json:"videoCount" binding:"required"`
	ViewCount       int `json:"viewCount" binding:"required"`
	SubscriberCount int `json:"subscriberCount" binding:"required"`
}
