package main

import (
	"github.com/daariikk/MedNote/services/reminder-service/internal/app"
	"github.com/daariikk/MedNote/services/reminder-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	application := app.New(cfg)
	application.Run()
}
