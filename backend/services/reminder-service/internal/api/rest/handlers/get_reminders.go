package handlers

import (
	"fmt"
	"github.com/daariikk/MedNote/services/reminder-service/internal/api/response"
	"github.com/daariikk/MedNote/services/reminder-service/internal/domain"
	"log/slog"
	"net/http"
	"strconv"
)

type GetReminderWrapper interface {
	GetReminders(userId int64) ([]domain.Reminder, error)
}

func GetRemindersHandler(logger *slog.Logger, reminder GetReminderWrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("Запущен GetRemindersHandler для получения всех напоминаний")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		const op = "record-service/internal/api/rest/handlers/get_reminders/GetRemindersHandler"
		logger.With(op)

		userIdStr := r.URL.Query().Get("userId")

		if userIdStr == "" {
			http.Error(w, "Missing required parameters", http.StatusBadRequest)
			return
		}

		logger.Debug("Query-параметры успешно извлечены")

		// Преобразование userId в число
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid userId", http.StatusBadRequest)
			return
		}

		// Логика получения записей
		reminders, err := reminder.GetReminders(userId)
		if err != nil {
			response.SendFailureResponse(w, fmt.Sprintf("Failed to get reminders: %v", err), http.StatusInternalServerError)
			return
		}

		logger.Info("Все напоминания успешно получены и переданы", "userId", userId)
		response.SendSuccessResponse(w, reminders, http.StatusOK)
	}
}
