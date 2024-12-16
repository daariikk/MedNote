package postgres

import (
	"context"
	"fmt"
	"github.com/daariikk/MedNote/services/reminder-service/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log/slog"
)

type Storage struct {
	connection *pgx.Conn
	logger     *slog.Logger
}

func New(ctx context.Context, logger *slog.Logger, url string) (*Storage, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		logger.Error("Failed to connect to postgres", "error", err)
		return nil, errors.Wrapf(err, "failed to connect to postgres")
	}

	return &Storage{conn, logger}, nil
}

func (s *Storage) Close() error {
	if s.connection != nil {
		return s.connection.Close(context.Background())
	}
	return nil
}

func (s *Storage) DeleteReminder(patientId int64, reminderId int64) error {
	query := fmt.Sprintf("DELETE FROM reminders WHERE id=$1 AND patient_id=$2")
	tag, err := s.connection.Exec(context.Background(), query, reminderId, patientId)
	if err != nil {
		return errors.Wrap(err, "failed to execute delete query")
	}
	if tag.RowsAffected() == 0 {

		return repository.ErrorDeleteFailed
	}
	return nil
}
