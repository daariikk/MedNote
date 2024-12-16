package postgres

import (
	"context"
	"fmt"
	"github.com/daariikk/MedNote/services/record-service/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log/slog"
)

type Storage struct {
	connection *pgx.Conn
	logger     *slog.Logger
}

type Record struct {
	Type           string `json:"type"`
	ID             int64  `json:"id"`
	UpperIndicator int64  `json:"upper_indicator,omitempty"`
	LowerIndicator int64  `json:"lower_indicator,omitempty"`
	Indicator      int64  `json:"indicator,omitempty"`
	Hours          int64  `json:"hours,omitempty"`
	Minutes        int64  `json:"minutes,omitempty"`
	VolumeGlass    int64  `json:"volume_glass,omitempty"`
	CountGlass     int64  `json:"count_glass,omitempty"`
	Control        string `json:"control,omitempty"`
	DateOfAddition string `json:"date_of_addition"`
	IdPatient      int64  `json:"patient_id"`
	Complaint      string `json:"complaint,omitempty"`
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

func (s *Storage) DeleteRecords(patientId int64, tableName string, recordId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND patient_id=$2", tableName)
	tag, err := s.connection.Exec(context.Background(), query, recordId, patientId)
	if err != nil {
		return errors.Wrap(err, "failed to execute delete query")
	}
	if tag.RowsAffected() == 0 {

		return repository.ErrorDeleteFailed
	}
	return nil
}
