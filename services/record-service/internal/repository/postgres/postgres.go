package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/patient-service/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log/slog"
)

type Storage struct {
	connection *pgx.Conn
}

func New(ctx context.Context, logger *slog.Logger, url string) (*Storage, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		logger.Error("Failed to connect to postgres", "error", err)
		return nil, errors.Wrapf(err, "failed to connect to postgres")
	}

	return &Storage{conn}, nil
}

func (s *Storage) Close() error {
	if s.connection != nil {
		return s.connection.Close(context.Background())
	}
	return nil
}

func (s *Storage) GetPatient(userId int64) (domain.Patient, error) {
	var patient domain.Patient
	err := s.connection.QueryRow(context.Background(), `
        SELECT id, first_name, second_name, email, height, weight, gender, password, registration_data
        FROM patients
        WHERE id = $1
    `, userId).Scan(
		&patient.Id,
		&patient.FirstName,
		&patient.SecondName,
		&patient.Email,
		&patient.Height,
		&patient.Weight,
		&patient.Gender,
		&patient.Password,
		&patient.RegistrationData,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Patient{}, errors.Wrapf(err, "patient with id %d not found", userId)
		}
		return domain.Patient{}, errors.Wrapf(err, "failed to get patient with id %d", userId)
	}
	return patient, nil
}

func (s *Storage) UpdatePatient(patient domain.Patient) (domain.Patient, error) {
	_, err := s.connection.Exec(context.Background(), `
        UPDATE patients
        SET first_name = $1, second_name = $2, email = $3, height = $4, weight = $5, gender = $6, password = $7, registration_data = $8
        WHERE id = $9
    `, patient.FirstName, patient.SecondName, patient.Email, patient.Height, patient.Weight, patient.Gender, patient.Password, patient.RegistrationData, patient.Id)
	if err != nil {
		return domain.Patient{}, errors.Wrapf(err, "failed to update patient with id %d", patient.Id)
	}
	return patient, nil
}
