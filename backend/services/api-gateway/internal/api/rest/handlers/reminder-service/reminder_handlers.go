package reminder_service

import (
	"fmt"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"log/slog"
	"net/http"
)

func NewReminder(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/api/v1/reminders", cfg.Services.ReminderService)
		handlers.ForwardRequest(logger, w, r, url, "POST")
	}
}

func GetReminders(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/api/v1/reminders", cfg.Services.ReminderService)
		handlers.ForwardRequest(logger, w, r, url, "GET")
	}
}

func DeleteReminder(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/api/v1/reminders", cfg.Services.ReminderService)
		handlers.ForwardRequest(logger, w, r, url, "DELETE")
	}
}
