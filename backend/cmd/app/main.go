package main

import (
	"log"

	"github.com/korpgoodness/service.git/internal/app"
)

func main() {
	server := new(app.Server)
	if err := server.Run(); err != nil {
		log.Fatalf("Error run server %s", err.Error())
	}
}
