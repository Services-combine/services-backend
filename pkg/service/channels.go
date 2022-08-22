package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const (
	SCRIPT_GET_DATA_CHANNEL = "get_data_channel.py"
)

type ChannelsService struct {
	repo repository.Channels
}

func NewChannelsService(repo repository.Channels) *ChannelsService {
	return &ChannelsService{repo: repo}
}

func (s *ChannelsService) Add(ctx context.Context, channel domain.ChannelAdd) error {
	channelData, err := s.GetChannelById(ctx, channel.ChannelId, channel.ApiKey)
	if err != nil {
		return err
	}

	err = s.repo.Add(ctx, channelData)
	return err
}

func (s *ChannelsService) GetChannelById(ctx context.Context, channelId, apiKey string) (domain.ChannelAdd, error) {
	channel := domain.ChannelAdd{
		ChannelId: channelId,
		ApiKey:    apiKey,
	}

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return domain.ChannelAdd{}, err
	}

	call := youtubeService.Channels.List([]string{"snippet", "statistics"})
	call.Id(channelId)
	response, err := call.Do()
	if err != nil {
		return domain.ChannelAdd{}, err
	}

	channel.Title = response.Items[0].Snippet.Title
	channel.Description = response.Items[0].Snippet.Description
	channel.Photo = response.Items[0].Snippet.Thumbnails.Default.Url
	channel.ViewCount = response.Items[0].Statistics.ViewCount
	channel.SubscriberCount = response.Items[0].Statistics.SubscriberCount
	channel.VideoCount = response.Items[0].Statistics.VideoCount

	return channel, nil
}
