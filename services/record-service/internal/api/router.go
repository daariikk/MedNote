package api

import (
	"github.com/daariikk/MedNote/services/record-service/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/record-service/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func NewRouter(logger *slog.Logger, storage *postgres.Storage) *chi.Mux {
	router := chi.NewRouter()

	router.Post("/api/v1/records/{tableName}", handlers.NewRecord(logger, storage))
	router.Get("/api/v1/records", handlers.GetRecords(logger, storage))
	router.Delete("/api/v1/records", handlers.DeleteRecord(logger, storage))
	// router.Put("/api/v1/records/{recordId}", handlers.UpdateRecord(logger, storage))

	return router
}
