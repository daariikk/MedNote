package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/record-service/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log"
)

func (s *Storage) GetAllRecords(userId int64, date string) ([]Record, error) {
	tensions, err := s.GetTensions(userId, date)
	if err != nil {
		log.Println(err)
	}

	pulses, err := s.GetPulse(userId, date)
	if err != nil {
		log.Println(err)
	}

	complaints, err := s.GetComplaints(userId, date)
	if err != nil {
		log.Println(err)
	}

	steps, err := s.GetSteps(userId, date)
	if err != nil {
		log.Println(err)
	}

	sleeps, err := s.GetSleep(userId, date)
	if err != nil {
		log.Println(err)
	}

	waters, err := s.GetWater(userId, date)
	if err != nil {
		log.Println(err)
	}

	response := make([]Record, 0, 10)

	// Добавляем давление (tension)
	for _, tension := range tensions {
		response = append(response, Record{
			Type:           "tensions",
			ID:             tension.Id,
			UpperIndicator: tension.UpperIndicator,
			LowerIndicator: tension.LowerIndicator,
			Control:        tension.Control,
			DateOfAddition: tension.DateOfAddition,
			IdPatient:      tension.IdPatient,
		})
	}

	// Добавляем пульс (pulse)
	for _, pulse := range pulses {
		response = append(response, Record{
			Type:           "pulse",
			ID:             pulse.Id,
			Indicator:      pulse.Indicator,
			Control:        pulse.Control,
			DateOfAddition: pulse.DateOfAddition,
			IdPatient:      pulse.IdPatient,
		})
	}

	// Добавляем жалобы (complaints)
	for _, complaint := range complaints {
		response = append(response, Record{
			Type:           "complaints",
			ID:             complaint.Id,
			Complaint:      complaint.Complaint,
			DateOfAddition: complaint.DateOfAddition,
			IdPatient:      complaint.IdPatient,
		})
	}

	// Добавляем шаги (Steps)
	for _, step := range steps {
		response = append(response, Record{
			Type:           "steps",
			ID:             step.Id,
			Indicator:      step.Indicator,
			Control:        step.Control,
			DateOfAddition: step.DateOfAddition,
			IdPatient:      step.IdPatient,
		})
	}

	// Добавляем длительность сна (Sleep)
	for _, sleep := range sleeps {
		response = append(response, Record{
			Type:           "sleep",
			ID:             sleep.Id,
			Hours:          sleep.Hours,
			Minutes:        sleep.Minutes,
			Control:        sleep.Control,
			DateOfAddition: sleep.DateOfAddition,
			IdPatient:      sleep.IdPatient,
		})
	}

	// Добавляем контроль воды (Water)
	for _, water := range waters {
		response = append(response, Record{
			Type:           "water",
			ID:             water.Id,
			VolumeGlass:    water.VolumeGlass,
			CountGlass:     water.CountGlass,
			Indicator:      water.Indicator,
			Control:        water.Control,
			DateOfAddition: water.DateOfAddition,
			IdPatient:      water.IdPatient,
		})
	}
	return response, nil
}

func (s *Storage) GetPulse(userId int64, date string) ([]domain.Pulse, error) {
	records := make([]domain.Pulse, 0)
	query := `
SELECT * FROM pulse
WHERE patient_id=$1 AND date_of_addition=$2
`
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Pulse{}, errors.Wrapf(err, "pulse records not found")
		}
		return []domain.Pulse{}, errors.Wrapf(err, "failed to pulse records")
	}
	defer rows.Close()

	for rows.Next() {
		pulse := domain.Pulse{}
		err := rows.Scan(
			&pulse.Id,
			&pulse.Indicator,
			&pulse.Control,
			&pulse.DateOfAddition,
			&pulse.IdPatient,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, pulse)
	}
	return records, nil
}

func (s *Storage) GetSteps(userId int64, date string) ([]domain.Steps, error) {
	records := make([]domain.Steps, 0)
	query := `
SELECT * FROM steps
WHERE patient_id=$1 AND date_of_addition=$2
`
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Steps{}, errors.Wrapf(err, "steps records not found")
		}
		return []domain.Steps{}, errors.Wrapf(err, "failed to get steps records")
	}
	defer rows.Close()

	for rows.Next() {
		steps := domain.Steps{}
		err := rows.Scan(
			&steps.Id,
			&steps.Indicator,
			&steps.Control,
			&steps.DateOfAddition,
			&steps.IdPatient,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, steps)
	}
	return records, nil
}

func (s *Storage) GetTensions(userId int64, date string) ([]domain.Tensions, error) {
	records := make([]domain.Tensions, 0)
	query := `
SELECT * FROM tensions
WHERE patient_id=$1 AND date_of_addition=$2
`
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Tensions{}, errors.Wrapf(err, "tensions records not found")
		}
		return []domain.Tensions{}, errors.Wrapf(err, "failed to get tensions records")
	}
	defer rows.Close()

	for rows.Next() {
		tensions := domain.Tensions{}
		err := rows.Scan(
			&tensions.Id,
			&tensions.UpperIndicator,
			&tensions.LowerIndicator,
			&tensions.Control,
			&tensions.DateOfAddition,
			&tensions.IdPatient,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, tensions)
	}
	return records, nil
}

func (s *Storage) GetSleep(userId int64, date string) ([]domain.Sleep, error) {
	records := make([]domain.Sleep, 0)
	query := `
SELECT * FROM sleep
WHERE patient_id=$1 AND date_of_addition=$2
`
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Sleep{}, errors.Wrapf(err, "sleep records not found")
		}
		return []domain.Sleep{}, errors.Wrapf(err, "failed to get sleep records")
	}
	defer rows.Close()

	for rows.Next() {
		sleep := domain.Sleep{}
		err := rows.Scan(
			&sleep.Id,
			&sleep.Hours,
			&sleep.Minutes,
			&sleep.Control,
			&sleep.DateOfAddition,
			&sleep.IdPatient,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, sleep)
	}
	return records, nil
}

func (s *Storage) GetWater(userId int64, date string) ([]domain.Water, error) {
	records := make([]domain.Water, 0)
	query := `
SELECT * FROM water
WHERE patient_id=$1 AND date_of_addition=$2
`
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Water{}, errors.Wrapf(err, "water records not found")
		}
		return []domain.Water{}, errors.Wrapf(err, "failed to get water records")
	}
	defer rows.Close()

	for rows.Next() {
		water := domain.Water{}
		err := rows.Scan(
			&water.Id,
			&water.VolumeGlass,
			&water.CountGlass,
			&water.Indicator,
			&water.Control,
			&water.DateOfAddition,
			&water.IdPatient,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, water)
	}
	return records, nil
}

func (s *Storage) GetComplaints(userId int64, date string) ([]domain.Complaints, error) {
	records := make([]domain.Complaints, 0)
	query := `
SELECT * FROM complaints
WHERE patient_id=$1 AND date_of_addition=$2
`
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []domain.Complaints{}, errors.Wrapf(err, "complaints records not found")
		}
		return []domain.Complaints{}, errors.Wrapf(err, "failed to get complaints records")
	}
	defer rows.Close()

	for rows.Next() {
		complaints := domain.Complaints{}
		err := rows.Scan(
			&complaints.Id,
			&complaints.Complaint,
			&complaints.DateOfAddition,
			&complaints.IdPatient,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		records = append(records, complaints)
	}
	return records, nil
}
