package main

import (
	"github.com/daariikk/MedNote/services/patient-service/internal/app"
	"github.com/daariikk/MedNote/services/patient-service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	application := app.New(cfg)
	application.Run()
}
