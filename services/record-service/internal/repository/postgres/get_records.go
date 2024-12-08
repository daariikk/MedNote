package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/record-service/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (s *Storage) GetAllRecords(userId int64, date time.Time) ([]Record, error) {
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

// Общая функция для получения записей по таблице
func (s *Storage) getRecords(query string, userId int64, date time.Time, recordStruct interface{}) error {
	rows, err := s.connection.Query(context.Background(), query, userId, date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.Wrapf(err, "no records found")
		}
		return errors.Wrapf(err, "failed to get records")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(recordStruct)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}

func (s *Storage) GetPulse(userId int64, date time.Time) ([]domain.Pulse, error) {
	records := make([]domain.Pulse, 0, 5)
	query := `
SELECT * FROM pulse
WHERE patient_id = $1 AND date_of_addition = $2
`
	err := s.getRecords(query, userId, date, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *Storage) GetSteps(userId int64, date time.Time) ([]domain.Steps, error) {
	records := make([]domain.Steps, 0, 5)
	query := `
SELECT * FROM steps
WHERE patient_id = $1 AND date_of_addition = $2
`
	err := s.getRecords(query, userId, date, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
func (s *Storage) GetTensions(userId int64, date time.Time) ([]domain.Tensions, error) {
	records := make([]domain.Tensions, 0, 5)
	query := `
SELECT * FROM tensions
WHERE patient_id = $1 AND date_of_addition = $2
`
	err := s.getRecords(query, userId, date, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *Storage) GetSleep(userId int64, date time.Time) ([]domain.Sleep, error) {
	records := make([]domain.Sleep, 0, 5)
	query := `
SELECT * FROM sleep
WHERE patient_id = $1 AND date_of_addition = $2
`
	err := s.getRecords(query, userId, date, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *Storage) GetWater(userId int64, date time.Time) ([]domain.Water, error) {
	records := make([]domain.Water, 0, 5)
	query := `
SELECT * FROM water
WHERE patient_id = $1 AND date_of_addition = $2
`
	err := s.getRecords(query, userId, date, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (s *Storage) GetComplaints(userId int64, date time.Time) ([]domain.Complaints, error) {
	records := make([]domain.Complaints, 0, 5)
	query := `
SELECT * FROM complaints
WHERE patient_id = $1 AND date_of_addition = $2
`
	err := s.getRecords(query, userId, date, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
