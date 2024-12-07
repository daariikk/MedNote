package handlers

import (
	"encoding/json"
	"github.com/daariikk/MedNote/services/patient-service/internal/api/response"
	"github.com/daariikk/MedNote/services/patient-service/internal/domain"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"strconv"
)

type UpdatePatienter interface {
	UpdatePatient(patient domain.Patient) (domain.Patient, error)
}

func UpdatePatient(logger *slog.Logger, patienter UpdatePatienter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "patient-service/internal/api/rest/handlers/get_patient/Patient"
		//logger.With(op)

		userIdStr := chi.URLParam(r, "userId")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			response.SendFailureResponse(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		logger.Info("Handling UPDATE patient request for user", "userId", userId)

		var patient domain.Patient
		err = json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			response.SendFailureResponse(w, "Invalid request data", http.StatusBadRequest)
			return
		}

		logger.Debug("Handling UPDATE patient request for user", slog.Any("patient", patient))

		patient.Id = userId

		updatedPatient, err := patienter.UpdatePatient(patient)
		if err != nil {
			logger.Info("getting user error", slog.Any("error", err))
			response.SendFailureResponse(w, "Failed to update patient", http.StatusInternalServerError)
			return
		}

		logger.Info("Successful updating patient", "userId", userId)

		response.SendSuccessResponse(w, updatedPatient, http.StatusOK)
	}
}
