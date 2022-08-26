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

type ChannelsService struct {
	repo repository.Channels
}

func NewChannelsService(repo repository.Channels) *ChannelsService {
	return &ChannelsService{repo: repo}
}

func (s *ChannelsService) CheckingUniqueness(ctx context.Context, channel_id string) (bool, error) {
	status, err := s.repo.CheckingUniqueness(ctx, channel_id)
	return status, err
}

func (s *ChannelsService) Add(ctx context.Context, channel domain.ChannelAdd) error {
	channelApi, err := GetById(ctx, channel.ChannelId, channel.ApiKey)
	if err != nil {
		return err
	}

	channel.Title = channelApi.Title
	channel.Description = channelApi.Description
	channel.Photo = channelApi.Photo
	channel.ViewCount = channelApi.ViewCount
	channel.SubscriberCount = channelApi.SubscriberCount
	channel.VideoCount = channelApi.VideoCount
	channel.Launch = false
	channel.Comment = ""
	channel.CountCommentedVideos = 0

	err = s.repo.Add(ctx, channel)
	return err
}

func GetById(ctx context.Context, channelId, apiKey string) (domain.ChannelUpdate, error) {
	var channel domain.ChannelUpdate

	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return domain.ChannelUpdate{}, err
	}

	call := youtubeService.Channels.List([]string{"snippet", "statistics"})
	response, err := call.Id(channelId).Do()
	if err != nil {
		return domain.ChannelUpdate{}, domain.ErrInvalidApiKey
	}

	if response.Items == nil {
		return domain.ChannelUpdate{}, domain.ErrInvalidChannelId
	}

	channel.Title = response.Items[0].Snippet.Title
	channel.Description = response.Items[0].Snippet.Description
	channel.Photo = response.Items[0].Snippet.Thumbnails.Default.Url
	channel.ViewCount = response.Items[0].Statistics.ViewCount
	channel.SubscriberCount = response.Items[0].Statistics.SubscriberCount
	channel.VideoCount = response.Items[0].Statistics.VideoCount

	return channel, nil
}

func (s *ChannelsService) Get(ctx context.Context) ([]domain.ChannelGet, error) {
	channels, err := s.repo.Get(ctx)
	return channels, err
}

func (s *ChannelsService) Launch(ctx context.Context, channelID primitive.ObjectID) error {
	err := s.repo.Launch(ctx, channelID)
	return err
}

func (s *ChannelsService) Update(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelIdKey) error {
	channelData, err := GetById(ctx, channel.ChannelId, channel.ApiKey)
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, channelID, channelData)
	return err
}

func (s *ChannelsService) Delete(ctx context.Context, channelID primitive.ObjectID, channel_id string) error {
	if err := s.repo.Delete(ctx, channelID); err != nil {
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

func (s *ChannelsService) EditProxy(ctx context.Context, channelID primitive.ObjectID, channel domain.ProxyEdit) error {
	err := s.repo.EditProxy(ctx, channelID, channel)
	return err
}
