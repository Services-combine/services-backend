package repository

import (
	"github.com/b0shka/services/pkg/database/mongodb"
	"github.com/b0shka/services/pkg/logger"
	"os"
	"testing"
)

var testRepos *Repositories

func TestMain(m *testing.M) {
	mongodbURL := "mongodb://127.0.0.1:27017"

	mongoClient, err := mongodb.NewClient(mongodbURL)
	if err != nil {
		logger.Errorf("Error connect mongodb: %s", err.Error())
	}

	testRepos = NewRepositories(mongoClient, "services")

	os.Exit(m.Run())
}
