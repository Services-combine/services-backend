package server

import (
	"context"
	"net/http"

	"github.com/b0shka/services/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:           ":" + cfg.HTTP.Port,
		Handler:        handler,
		ReadTimeout:    cfg.HTTP.ReadTimeout,
		WriteTimeout:   cfg.HTTP.WriteTimeout,
		MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
	}

	server := &Server{
		httpServer: httpServer,
	}

	return server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
