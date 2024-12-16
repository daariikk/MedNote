package handlers

import (
	"errors"
	"github.com/daariikk/MedNote/services/patient-service/internal/api/response"
	"github.com/daariikk/MedNote/services/patient-service/internal/domain"
	"github.com/daariikk/MedNote/services/patient-service/internal/lib/logger/sl"
	"github.com/daariikk/MedNote/services/patient-service/internal/repository"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strconv"
)

type GetPatienter interface {
	GetPatient(userId int64) (domain.Patient, error)
}

func Patient(logger *slog.Logger, patienter GetPatienter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "patient-service/internal/api/rest/handlers/get_patient/Patient"
		logger.With(op)

		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			response.SendFailureResponse(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		logger.Info("Handling GET patient request for user", "userId", userId)

		patient, err := patienter.GetPatient(userId)
		if err != nil {
			if errors.Is(err, repository.ErrorNotFound) {
				logger.Debug("Patient not found", sl.Err(err))
				response.SendFailureResponse(w, "Patient not found", http.StatusNotFound)
			} else {
				logger.Debug("Patient not found", sl.Err(err))
				response.SendFailureResponse(w, "Failed to get patient", http.StatusInternalServerError)
			}
			return
		}

		logger.Info("Successful getting patient from storage", "userId", userId)

		response.SendSuccessResponse(w, patient, http.StatusOK)
	}
}
