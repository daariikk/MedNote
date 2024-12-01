package api

import (
	"github.com/daariikk/MedNote/services/patient-service/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/patient-service/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"log/slog"
)

func NewRouter(logger *slog.Logger, storage *postgres.Storage) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/api/v1/users/{userId}/patient", handlers.Patient(logger, storage))
	router.Put("/api/v1/users/{userId}/patient", handlers.UpdatePatient(logger, storage))

	return router
}
