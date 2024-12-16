package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/reminder-service/internal/domain"
	"github.com/pkg/errors"
)

func (s *Storage) CreateReminder(reminder domain.Reminder) (domain.Reminder, error) {
	query := `
        INSERT INTO reminders (title, text, date, time, patient_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		reminder.Title,
		reminder.Text,
		reminder.Date,
		reminder.Time,
		reminder.PatientId,
	).Scan(&id)
	if err != nil {
		return domain.Reminder{}, errors.Wrapf(err, "failed to insert reminder")
	}
	reminder.Id = id
	return reminder, nil
}
