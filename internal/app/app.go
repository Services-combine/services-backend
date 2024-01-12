package app

import (
	"context"
	"errors"
	"github.com/b0shka/services/internal/config"
	"github.com/b0shka/services/internal/handler"
	"github.com/b0shka/services/internal/repository"
	"github.com/b0shka/services/internal/server"
	"github.com/b0shka/services/internal/service"
	"github.com/b0shka/services/internal/worker"
	"github.com/b0shka/services/pkg/auth"
	"github.com/b0shka/services/pkg/database/mongodb"
	"github.com/b0shka/services/pkg/email"
	"github.com/b0shka/services/pkg/hash"
	"github.com/b0shka/services/pkg/identity"
	"github.com/b0shka/services/pkg/logger"
	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			Services API
//	@version		1.0
//	@description	REST API for Services App

//	@host		localhost:8080
//	@BasePath	/api/v1/

//	@securityDefinitions.apikey	UsersAuth
//	@in							header
//	@name						Authorization

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logger.Error(err)

		return
	}

	hasher, err := hash.NewSHA256Hasher(cfg.Auth.CodeSalt)
	if err != nil {
		logger.Error(err)

		return
	}

	tokenManager, err := auth.NewPasetoManager(cfg.Auth.SecretKey)
	if err != nil {
		logger.Error(err)

		return
	}

	idGenerator := identity.NewIDGenerator()

	mongoClient, err := mongodb.NewClient(cfg.Mongo.URL)
	if err != nil {
		logger.Errorf("Error connect mongodb: %s", err.Error())

		return
	}

	logger.Info("Success connect mongodb")

	repos := repository.NewRepositories(mongoClient, cfg.Mongo.DatabaseName)

	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.Redis.Address,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	go runTaskProcessor(redisOpt, repos, hasher, idGenerator, cfg)

	services := service.NewServices(service.Deps{
		Repos:           repos,
		Hasher:          hasher,
		TokenManager:    tokenManager,
		AuthConfig:      cfg.Auth,
		TaskDistributor: taskDistributor,
		FoldersConfig:   cfg.Folders,
	})

	handlers := handler.NewHandler(
		services,
		tokenManager,
		cfg.Folders,
	)
	routes := handlers.InitRoutes(cfg)
	srv := server.NewServer(cfg, routes)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")
	gracefulShutdown(srv, mongoClient)
}

func gracefulShutdown(srv *server.Server, mongoClient *mongo.Client) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)

	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("Failed to stop server: %v", err)
	}

	logger.Info("Server stopped")

	if err := mongoClient.Disconnect(ctx); err != nil {
		logger.Errorf("Failed to disconnect mongodb: %v", err)
	}

	logger.Info("Mongodb disconnected")
}

func runTaskProcessor(
	redisOpt asynq.RedisClientOpt,
	repos *repository.Repositories,
	hasher hash.Hasher,
	idGenerator identity.Generator,
	cfg *config.Config,
) {
	emailService := email.NewEmailService(
		cfg.Email.ServiceName,
		cfg.Email.ServiceAddress,
		cfg.Email.ServicePassword,
		cfg.SMTP.Host,
		cfg.SMTP.Port,
	)

	taskProcessor := worker.NewRedisTaskProcessor(
		redisOpt,
		repos,
		hasher,
		idGenerator,
		emailService,
		cfg.Email,
		cfg.Auth,
	)

	logger.Info("Start task processor")

	if err := taskProcessor.Start(); err != nil {
		logger.Error("Failed to start task processor")
	}
}
