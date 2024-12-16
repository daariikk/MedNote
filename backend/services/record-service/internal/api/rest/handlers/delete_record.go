package handlers

import (
	"errors"
	"fmt"
	"github.com/daariikk/MedNote/services/record-service/internal/api/response"
	"github.com/daariikk/MedNote/services/record-service/internal/repository"
	"log/slog"
	"net/http"
	"strconv"
)

type DeleteRecorder interface {
	DeleteRecords(patientId int64, tableName string, recordId int64) error
}

func DeleteRecord(logger *slog.Logger, recorder DeleteRecorder) http.HandlerFunc {
	logger.Debug("DeleteRecord start...")
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "record-service/internal/api/rest/handlers/delete_record/DeleteRecord"
		logger.With(op)
		logger.Debug("URL", slog.String("url", r.URL.String()))
		userIdStr := r.URL.Query().Get("userId")
		tableName := r.URL.Query().Get("tableName")
		recordIdStr := r.URL.Query().Get("recordId")
		logger.Debug("userIdStr", slog.String("userIdStr", userIdStr))
		logger.Debug("tableName", slog.String("tableName", userIdStr))
		logger.Debug("recordId", slog.String("recordId", userIdStr))
		// Преобразование userId в число
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			response.SendFailureResponse(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Преобразование recordId в число
		recordId, err := strconv.ParseInt(recordIdStr, 10, 64)
		if err != nil {
			response.SendFailureResponse(w, "Invalid record ID", http.StatusBadRequest)
			return
		}

		// Логика удаления записи
		err = recorder.DeleteRecords(userId, tableName, recordId)
		if err != nil {
			if errors.Is(err, repository.ErrorDeleteFailed) {
				response.SendFailureResponse(w, fmt.Sprintf("Failed to delete record: %v", err), http.StatusNotFound)
			} else {
				response.SendFailureResponse(w, fmt.Sprintf("Failed to delete record: %v", err), http.StatusInternalServerError)
			}
			return
		}

		logger.Info("Handling DELETE record request for user", "userId", userId)
		response.SendSuccessResponse(w, "Record deleted successfully", http.StatusNoContent)
	}
}
