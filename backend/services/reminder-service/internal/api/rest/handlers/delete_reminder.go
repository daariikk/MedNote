package handlers

import (
	"errors"
	"fmt"
	"github.com/daariikk/MedNote/services/reminder-service/internal/api/response"
	"github.com/daariikk/MedNote/services/reminder-service/internal/repository"
	"log/slog"
	"net/http"
	"strconv"
)

type DeleteReminderWrapper interface {
	DeleteReminder(patientId int64, recordId int64) error
}

func DeleteReminderHandler(logger *slog.Logger, reminder DeleteReminderWrapper) http.HandlerFunc {
	logger.Debug("DeleteReminder start...")
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "record-service/internal/api/rest/handlers/delete_reminder/DeleteReminderHandler"
		logger.With(op)
		logger.Debug("URL", slog.String("url", r.URL.String()))
		userIdStr := r.URL.Query().Get("userId")
		reminderIdStr := r.URL.Query().Get("reminderId")
		logger.Debug("userIdStr", slog.String("userIdStr", userIdStr))
		logger.Debug("reminderIdStr", slog.String("reminderIdStr", reminderIdStr))
		// Преобразование userId в число
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			response.SendFailureResponse(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Преобразование recordId в число
		reminderId, err := strconv.ParseInt(reminderIdStr, 10, 64)
		if err != nil {
			response.SendFailureResponse(w, "Invalid reminder ID", http.StatusBadRequest)
			return
		}

		// Логика удаления записи
		err = reminder.DeleteReminder(userId, reminderId)
		if err != nil {
			if errors.Is(err, repository.ErrorDeleteFailed) {
				response.SendFailureResponse(w, fmt.Sprintf("Failed to delete reminder: %v", err), http.StatusNotFound)
			} else {
				response.SendFailureResponse(w, fmt.Sprintf("Failed to delete reminder: %v", err), http.StatusInternalServerError)
			}
			return
		}

		logger.Info("Handling DELETE reminder request for user", "userId", userId)
		response.SendSuccessResponse(w, "Reminder deleted successfully", http.StatusNoContent)
	}
}
