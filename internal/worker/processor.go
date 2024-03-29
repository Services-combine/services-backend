package worker

import (
	"context"
	"encoding/json"

	"github.com/b0shka/services/internal/config"
	"github.com/b0shka/services/internal/repository"
	"github.com/b0shka/services/pkg/email"
	"github.com/b0shka/services/pkg/hash"
	"github.com/b0shka/services/pkg/identity"
	"github.com/b0shka/services/pkg/logger"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendLoginNotification(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server       *asynq.Server
	repos        *repository.Repositories
	hasher       hash.Hasher
	idGenerator  identity.Generator
	emailService *email.EmailService
	emailConfig  config.EmailConfig
	authConfig   config.AuthConfig
}

func NewRedisTaskProcessor(
	redisOpt asynq.RedisClientOpt,
	repos *repository.Repositories,
	hasher hash.Hasher,
	idGenerator identity.Generator,
	emailService *email.EmailService,
	emailConfig config.EmailConfig,
	authConfig config.AuthConfig,
) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				var data map[string]interface{}
				verr := json.Unmarshal(task.Payload(), &data)
				if verr != nil {
					logger.Errorf("Error decode payload: %s", verr.Error())

					return
				}

				logger.Errorf("process task failed: type - %s, payload - %v, err - %s", task.Type(), data, err.Error())
			}),
			// Logger: logger,
		},
	)

	return &RedisTaskProcessor{
		server:       server,
		repos:        repos,
		hasher:       hasher,
		idGenerator:  idGenerator,
		emailService: emailService,
		emailConfig:  emailConfig,
		authConfig:   authConfig,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendLoginNotification, processor.ProcessTaskSendLoginNotification)

	return processor.server.Start(mux)
}
