package api

import (
	"github.com/daariikk/MedNote/services/patient-service/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/patient-service/internal/repository/postgres"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func NewRouter(logger *slog.Logger, storage *postgres.Storage) *chi.Mux {
	router := chi.NewRouter()
	router.Options("/api/v1/users/{userId}/patient", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
	})
	router.Get("/api/v1/users/{userId}/patient", handlers.Patient(logger, storage))
	router.Put("/api/v1/users/{userId}/patient", handlers.UpdatePatient(logger, storage))

	return router
}
