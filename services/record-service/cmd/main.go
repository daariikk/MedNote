package main

import (
	"github.com/daariikk/MedNote/services/record-service/internal/app"
	"github.com/daariikk/MedNote/services/record-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	application := app.New(cfg)
	application.Run()
}
