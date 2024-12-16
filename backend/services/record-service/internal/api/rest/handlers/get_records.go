package handlers

import (
	"fmt"
	"github.com/daariikk/MedNote/services/record-service/internal/api/response"
	"github.com/daariikk/MedNote/services/record-service/internal/repository/postgres"
	"log/slog"
	"net/http"
	"strconv"
)

type GetRecorder interface {
	GetAllRecords(userId int64, date string) ([]postgres.Record, error)
}

func GetRecords(logger *slog.Logger, recorder GetRecorder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("GetRecords starting...")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		const op = "record-service/internal/api/rest/handlers/get_records/GetRecords"
		logger.With(op)

		userIdStr := r.URL.Query().Get("userId")
		date := r.URL.Query().Get("date")

		logger.Debug("Check UserId and Date", slog.String("userId", userIdStr), slog.String("date", date))

		if userIdStr == "" || date == "" {
			http.Error(w, "Missing required parameters", http.StatusBadRequest)
			return
		}

		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid userId", http.StatusBadRequest)
			return
		}

		records, err := recorder.GetAllRecords(userId, date)
		if err != nil {
			response.SendFailureResponse(w, fmt.Sprintf("Failed to get records: %v", err), http.StatusInternalServerError)
			return
		}

		logger.Info("Handling GET All records request for user", "userId", userId, "date", date)
		logger.Debug("GetRecords works successful")
		response.SendSuccessResponse(w, records, http.StatusOK)
	}
}
