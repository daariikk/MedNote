package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/patient-service/internal/domain"
	"github.com/daariikk/MedNote/services/patient-service/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Storage struct {
	connectionPool *pgxpool.Pool
}

func New(ctx context.Context, url string) (*Storage, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to postgres")
	}

	return &Storage{pool}, nil
}

func (s *Storage) Close() error {
	if s.connectionPool != nil {
		s.connectionPool.Close()
	}
	return nil
}

func (s *Storage) GetPatient(userId int64) (domain.Patient, error) {
	var patient domain.Patient
	err := s.connectionPool.QueryRow(context.Background(), `
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
			return domain.Patient{}, repository.ErrorNotFound
		}
		return domain.Patient{}, errors.Wrapf(err, "failed to get patient with id %d", userId)
	}
	return patient, nil
}

func (s *Storage) UpdatePatient(patient domain.Patient) (domain.Patient, error) {
	_, err := s.connectionPool.Exec(context.Background(), `
        UPDATE patients
        SET first_name = $1, second_name = $2, height = $3, weight = $4, gender = $5
        WHERE id = $6
    `, patient.FirstName, patient.SecondName, patient.Height, patient.Weight, patient.Gender, patient.Id)
	if err != nil {
		return domain.Patient{}, errors.Wrapf(err, "failed to update patient with id %d", patient.Id)
	}
	return patient, nil
}
