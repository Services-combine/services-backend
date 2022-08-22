package service

import (
	"context"
	"os"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *ChannelsService) Add(ctx context.Context, channel domain.ChannelIdKey) error {
	channelData, err := GetChannelById(ctx, channel.ChannelId, channel.ApiKey)
	if err != nil {
		return err
	}

	err = s.repo.Add(ctx, channelData)
	return err
}

func GetChannelById(ctx context.Context, channelId, apiKey string) (domain.ChannelAdd, error) {
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
	channel.Launch = false

	return channel, nil
}

func (s *ChannelsService) GetChannels(ctx context.Context) ([]domain.ChannelGet, error) {
	channels, err := s.repo.GetChannels(ctx)
	return channels, err
}

func (s *ChannelsService) LaunchChannel(ctx context.Context, channelID primitive.ObjectID) error {
	err := s.repo.LaunchChannel(ctx, channelID)
	return err
}

func (s *ChannelsService) UpdateChannel(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelIdKey) error {
	channelData, err := GetChannelById(ctx, channel.ChannelId, channel.ApiKey)
	if err != nil {
		return err
	}

	err = s.repo.UpdateChannel(ctx, channelID, channelData)
	return err
}

func (s *ChannelsService) DeleteChannel(ctx context.Context, channelID primitive.ObjectID, channel_id string) error {
	if err := s.repo.DeleteChannel(ctx, channelID); err != nil {
		return err
	}

	err := os.Remove(os.Getenv("FOLDER_CHANNELS") + "app_token_" + channel_id + ".json")
	if err != nil {
		return err
	}

	err = os.Remove(os.Getenv("FOLDER_CHANNELS") + "user_token_" + channel_id + ".json")
	return nil
}

func (s *ChannelsService) EditChannel(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelEdit) error {
	err := s.repo.EditChannel(ctx, channelID, channel)
	return err
}
