package handlers

import (
	"fmt"
	"github.com/daariikk/MedNote/services/record-service/internal/api/response"
	"github.com/daariikk/MedNote/services/record-service/internal/repository/postgres"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type GetRecorder interface {
	GetAllRecords(userId int64, date time.Time) ([]postgres.Record, error)
}

func GetRecords(logger *slog.Logger, recorder GetRecorder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "record-service/internal/api/rest/handlers/get_records/GetRecords"
		logger.With(op)

		userIdStr := r.URL.Query().Get("userId")
		dateStr := r.URL.Query().Get("date")

		if userIdStr == "" || dateStr == "" {
			http.Error(w, "Missing required parameters", http.StatusBadRequest)
			return
		}

		// Преобразование userId в число
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid userId", http.StatusBadRequest)
			return
		}

		// Преобразование date в time.Time
		date, err := time.Parse("02.12.2024", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format. Expected dd.MM.yyyy", http.StatusBadRequest)
			return
		}

		// Логика получения записей
		records, err := recorder.GetAllRecords(userId, date)
		if err != nil {
			response.SendFailureResponse(w, fmt.Sprintf("Failed to get records: %v", err), http.StatusInternalServerError)
			return
		}

		logger.Info("Handling GET All records request for user", "userId", userId, "date", date)
		response.SendSuccessResponse(w, records, http.StatusOK)
	}
}
