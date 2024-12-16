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
	logger     *slog.Logger
}

func New(ctx context.Context, logger *slog.Logger, url string) (*Storage, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		logger.ErrorContext(ctx, "Failed to connect to postgres", "error", err)
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

func (s *Storage) GetPatient(userId int64) (domain.Patient, error) {
	var patient domain.Patient
	if s.connection == nil {
		return domain.Patient{}, errors.New("database connection is not initialized")
	}

	s.logger.Debug("init getting process patient", slog.Int64("user_id", userId))

	err := s.connection.QueryRow(context.Background(), `
        SELECT *
        FROM patients
        WHERE patient_id=$1
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
	s.logger.Debug("patient", patient)
	s.logger.Debug("err", err)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logger.Debug("patient not found in database")
			return domain.Patient{}, errors.Wrapf(err, "patient with id %d not found", userId)
		}
		s.logger.Debug("Я попал сюда?????????")
		return domain.Patient{}, errors.Wrapf(err, "failed to get patient with id %d", userId)
	}
	return patient, nil
}

func (s *Storage) UpdatePatient(patient domain.Patient) (domain.Patient, error) {
	_, err := s.connection.Exec(context.Background(), `
        UPDATE patients
        SET first_name=$1, second_name=$2, height=$3, weight=$4, gender=$5
        WHERE patient_id=$6
    `, patient.FirstName, patient.SecondName, patient.Height, patient.Weight, patient.Gender, patient.Id)

	err = s.connection.QueryRow(context.Background(), `
        SELECT patient_id, first_name, second_name, email, height, weight, gender, password, registration_date
        FROM patients
        WHERE patient_id = $1
    `, patient.Id).Scan(
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
		return domain.Patient{}, errors.Wrapf(err, "failed to update patient with id %d", patient.Id)
	}

	return patient, nil
}
