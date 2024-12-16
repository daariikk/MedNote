package main

import (
	"github.com/daariikk/MedNote/services/api-gateway/internal/app"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
)

func main() {
	cfg := config.MustLoad()
	application := app.New(cfg)
	application.Run()
}
