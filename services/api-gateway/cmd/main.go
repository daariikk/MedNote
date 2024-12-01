package main

import "github.com/daariikk/MedNote/services/patient-service/internal/config"

func main() {
	// Выгружаем конфиг
	cfg := config.MustLoad()
	// TODO: сделать миграции
}
