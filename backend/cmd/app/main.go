package main

import (
	"github.com/korpgoodness/service.git/internal/app"
	"github.com/korpgoodness/service.git/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	server := new(app.Server)
	if err := server.Run(); err != nil {
		logger.Errorf("Error run server %s", err.Error())
	}
}
