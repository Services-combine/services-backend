package main

import (
	"github.com/b0shka/services/internal/app"
)

const configPath = "configs"

func main() {
	app.Run(configPath)
}
