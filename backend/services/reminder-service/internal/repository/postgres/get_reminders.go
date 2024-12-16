package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/reminder-service/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log"
)

func (s *Storage) GetReminders(userId int64) ([]domain.Reminder, error) {
	reminders := make([]domain.Reminder, 0)
	query := `
SELECT * FROM reminders
WHERE patient_id=$1
`
	rows, err := s.connection.Query(context.Background(), query, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Reminder{}, errors.Wrapf(err, "reminder not found")
		}
		return []domain.Reminder{}, errors.Wrapf(err, "failed to reminder")
	}
	defer rows.Close()

	for rows.Next() {
		reminder := domain.Reminder{}
		err := rows.Scan(
			&reminder.Id,
			&reminder.Title,
			&reminder.Text,
			&reminder.Date,
			&reminder.Time,
			&reminder.PatientId,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		reminders = append(reminders, reminder)
	}
	return reminders, nil
}
