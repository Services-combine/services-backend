package app

import (
	"context"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/korpgoodness/service.git/pkg/handler"
	"github.com/korpgoodness/service.git/pkg/repository"
	"github.com/korpgoodness/service.git/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer *http.Server
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (s *Server) Run() error {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewMongoDB(os.Getenv("MONDO_DB_URL"))
	if err != nil {
		logrus.Fatalf("error connect mongodb %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	routes := handlers.InitRoutes()

	s.httpServer = &http.Server{
		Addr:           ":" + viper.GetString("http.port"),
		Handler:        routes,
		MaxHeaderBytes: viper.GetInt("http.maxHeaderBytes"),
		ReadTimeout:    viper.GetDuration("http.readTimeout"),
		WriteTimeout:   viper.GetDuration("http.writeTimeout"),
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
