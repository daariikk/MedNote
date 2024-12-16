package api

import (
	"github.com/daariikk/MedNote/services/reminder-service/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/reminder-service/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger, storage *postgres.Storage) *chi.Mux {
	router := chi.NewRouter()
	router.Options("/api/v1/reminders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})
	router.Post("/api/v1/reminders", handlers.NewReminderHandler(logger, storage))
	router.Get("/api/v1/reminders", handlers.GetRemindersHandler(logger, storage))
	router.Delete("/api/v1/reminders", handlers.DeleteReminderHandler(logger, storage))

	return router
}
