package main

import (
	"fmt"
	"log"

	"github.com/korpgoodness/services.git/internal/app"
)

func main() {
	server := new(app.Server)
	if err := server.Run(); err != nil {
		log.Fatal(fmt.Sprintf("Error run server %s", err.Error()))
	}
}
