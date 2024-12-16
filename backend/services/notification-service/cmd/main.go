package main

import (
	"github.com/daariikk/MedNote/services/notification-service/internal/app"
	"github.com/daariikk/MedNote/services/notification-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	application := app.New(cfg)
	application.Run()
}
