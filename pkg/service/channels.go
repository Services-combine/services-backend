package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

const (
	SCRIPT_GET_DATA_CHANNEL = "get_data_channel.py"
)

type Test struct {
	Kind string
}

type ChannelsService struct {
	repo repository.Channels
}

func NewChannelsService(repo repository.Channels) *ChannelsService {
	return &ChannelsService{repo: repo}
}

func (s *ChannelsService) Add(ctx context.Context, channel domain.ChannelAdd) error {
	script := os.Getenv("FOLDER_PYTHON_SCRIPTS_CHANNELS") + SCRIPT_GET_DATA_CHANNEL
	args_channel_id := fmt.Sprintf("-C %s", channel.ChannelId)
	args_api_key := fmt.Sprintf("-A %s", channel.ApiKey)

	data_channel_json, err := exec.Command(path_python, script, args_channel_id, args_api_key).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(data_channel_json))
	if strings.Contains(string(data_channel_json), "ERROR") {
		return fmt.Errorf("Ошибка при получении данных канала")
	}

	var data_channel map[string]interface{}
	//var data_channel Test
	if err = json.Unmarshal(data_channel_json, &data_channel); err != nil {
		return err
	}
	fmt.Println(data_channel)

	err = s.repo.Add(ctx, channel)
	return err
}
