package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/notification-service/internal/repository"
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

func (s *Storage) UpdatePassword(email string, password string) error {
	s.logger.Debug("Updating password", "email", email, "password", password)

	query := `
	UPDATE patients
	SET password=$1
	WHERE email=$2
`
	_, err := s.connection.Exec(context.Background(), query, password, email)
	if err != nil {
		s.logger.Error("Failed to update password", "email", email, "error", err)
		return errors.Wrap(err, repository.ErrorNotFound.Error())
	}

	s.logger.Debug("Successfully updated password", "email", email, "password", password)
	return nil
}
