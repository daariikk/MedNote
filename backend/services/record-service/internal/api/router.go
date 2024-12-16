package api

import (
	"github.com/daariikk/MedNote/services/record-service/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/record-service/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger, storage *postgres.Storage) *chi.Mux {
	router := chi.NewRouter()
	router.Options("/api/v1/records", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})
	router.Post("/api/v1/records", handlers.NewRecord(logger, storage))
	router.Get("/api/v1/records", handlers.GetRecords(logger, storage))
	router.Delete("/api/v1/records", handlers.DeleteRecord(logger, storage))
	// router.Put("/api/v1/records/{recordId}", handlers.UpdateRecord(logger, storage))

	return router
}
