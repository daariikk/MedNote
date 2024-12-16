package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/daariikk/MedNote/services/record-service/internal/api/response"
	"github.com/daariikk/MedNote/services/record-service/internal/repository"
	"github.com/daariikk/MedNote/services/record-service/internal/repository/postgres"
	"log/slog"
	"net/http"
)

type CreateRecorder interface {
	CreateRecord(req postgres.Record) (interface{}, error)
}

func NewRecord(logger *slog.Logger, recorder CreateRecorder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "record-service/internal/api/rest/handlers/new_record/CreateRecord"
		logger.With(op)

		request := postgres.Record{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		newRecord, err := recorder.CreateRecord(request)
		if err != nil {
			if errors.Is(err, repository.ErrorTypeNotSupported) {
				response.SendFailureResponse(w, fmt.Sprintf("Failed to insert record: %v", err), http.StatusNotFound)
			} else {
				response.SendFailureResponse(w, fmt.Sprintf("Failed to insert record: %v", err), http.StatusInternalServerError)
			}
			return
		}
		logger.Info("Handling POST record request for user")
		response.SendSuccessResponse(w, newRecord, http.StatusCreated)
	}

}
