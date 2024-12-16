package patient_service

import (
	"fmt"
	"github.com/daariikk/MedNote/services/api-gateway/internal/api/rest/handlers"
	"github.com/daariikk/MedNote/services/api-gateway/internal/config"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func GetPatient(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		url := fmt.Sprintf("%s/api/v1/users/%s/patient", cfg.Services.PatientService, userId)
		handlers.ForwardRequest(logger, w, r, url, "GET")
	}
}

func UpdatePatient(logger *slog.Logger, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		url := fmt.Sprintf("%s/api/v1/users/%s/patient", cfg.Services.PatientService, userId)
		handlers.ForwardRequest(logger, w, r, url, "PUT")
	}
}
