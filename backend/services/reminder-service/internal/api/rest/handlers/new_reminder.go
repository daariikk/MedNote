package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/daariikk/MedNote/services/reminder-service/internal/api/response"
	"github.com/daariikk/MedNote/services/reminder-service/internal/domain"
	"log/slog"
	"net/http"
)

type CreateReminderWrapper interface {
	CreateReminder(reminder domain.Reminder) (domain.Reminder, error)
}

func NewReminderHandler(logger *slog.Logger, reminder CreateReminderWrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "record-service/internal/api/rest/handlers/new_reminder/NewReminderHandler"
		logger.With(op)

		request := domain.Reminder{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		newRecord, err := reminder.CreateReminder(request)
		if err != nil {
			response.SendFailureResponse(w, fmt.Sprintf("Failed to insert reminder: %v", err), http.StatusInternalServerError)
			return
		}
		logger.Info("Handling POST reminder request for user")
		response.SendSuccessResponse(w, newRecord, http.StatusCreated)
	}

}
