package record_service

import (
	"fmt"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"log/slog"
	"net/http"
)

func NewRecord(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	logger.Debug("Переадресация POST-запроса сервису медицинских записей")
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/api/v1/records", cfg.Services.RecordService)
		handlers.ForwardRequest(logger, w, r, url, "POST")
	}
}

func GetRecords(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/api/v1/records", cfg.Services.RecordService)
		handlers.ForwardRequest(logger, w, r, url, "GET")
	}
}

func DeleteRecord(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/api/v1/records", cfg.Services.RecordService)
		handlers.ForwardRequest(logger, w, r, url, "DELETE")
	}
}
