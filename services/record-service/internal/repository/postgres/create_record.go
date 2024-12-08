package postgres

import (
	"context"
	"github.com/daariikk/MedNote/services/record-service/internal/domain"
	"github.com/daariikk/MedNote/services/record-service/internal/repository"
	"github.com/pkg/errors"
)

func (s *Storage) CreateRecord(req Record) (interface{}, error) {
	switch req.Type {
	case "tensions":
		tensions := domain.Tensions{
			UpperIndicator: req.UpperIndicator,
			LowerIndicator: req.LowerIndicator,
			InfoField: domain.InfoField{
				Control:        req.Control,
				DateOfAddition: req.DateOfAddition,
				IdPatient:      req.IdPatient,
			},
		}
		record, err := s.CreateRecordTensions(tensions)
		if err != nil {
			return nil, err
		}
		return record, nil

	case "pulse":
		pulse := domain.Pulse{
			Indicator: req.Indicator,
			InfoField: domain.InfoField{
				Control:        req.Control,
				DateOfAddition: req.DateOfAddition,
				IdPatient:      req.IdPatient,
			},
		}
		record, err := s.CreateRecordPulse(pulse)
		if err != nil {
			return nil, err
		}
		return record, nil

	case "steps":
		steps := domain.Steps{
			Indicator: req.Indicator,
			InfoField: domain.InfoField{
				Control:        req.Control,
				DateOfAddition: req.DateOfAddition,
				IdPatient:      req.IdPatient,
			},
		}
		record, err := s.CreateRecordSteps(steps)
		if err != nil {
			return nil, err
		}
		return record, nil

	case "sleep":
		sleep := domain.Sleep{
			Hours:   req.Hours,
			Minutes: req.Minutes,
			InfoField: domain.InfoField{
				Control:        req.Control,
				DateOfAddition: req.DateOfAddition,
				IdPatient:      req.IdPatient,
			},
		}
		record, err := s.CreateRecordSleep(sleep)
		if err != nil {
			return nil, err
		}
		return record, nil

	case "water":
		water := domain.Water{
			VolumeGlass: req.VolumeGlass,
			CountGlass:  req.CountGlass,
			Indicator:   req.Indicator,
			InfoField: domain.InfoField{
				Control:        req.Control,
				DateOfAddition: req.DateOfAddition,
				IdPatient:      req.IdPatient,
			},
		}
		record, err := s.CreateRecordWater(water)
		if err != nil {
			return nil, err
		}
		return record, nil

	case "complaints":
		complaints := domain.Complaints{
			Complaint:      req.Complaint,
			DateOfAddition: req.DateOfAddition,
			IdPatient:      req.IdPatient,
		}
		record, err := s.CreateRecordComplaints(complaints)
		if err != nil {
			return nil, err
		}
		return record, nil

	default:
		return nil, repository.ErrorTypeNotSupported
	}

}

func (s *Storage) CreateRecordPulse(pulse domain.Pulse) (domain.Pulse, error) {
	query := `
        INSERT INTO pulse (indicator, control, date_of_addition, patient_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		pulse.Indicator,
		pulse.Control,
		pulse.DateOfAddition,
		pulse.IdPatient,
	).Scan(&id)
	if err != nil {
		return domain.Pulse{}, errors.Wrapf(err, "failed to insert pulse record")
	}
	pulse.Id = id
	return pulse, nil
}

func (s *Storage) CreateRecordSteps(steps domain.Steps) (domain.Steps, error) {
	query := `
        INSERT INTO steps (indicator, control, date_of_addition, patient_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		steps.Indicator,
		steps.Control,
		steps.DateOfAddition,
		steps.IdPatient,
	).Scan(&id)
	if err != nil {
		return domain.Steps{}, errors.Wrapf(err, "failed to insert steps record")
	}
	steps.Id = id
	return steps, nil
}

func (s *Storage) CreateRecordTensions(tensions domain.Tensions) (domain.Tensions, error) {
	query := `
        INSERT INTO tensions (upper_indicator, lower_indicator, control, date_of_addition, patient_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		tensions.UpperIndicator,
		tensions.LowerIndicator,
		tensions.Control,
		tensions.DateOfAddition,
		tensions.IdPatient,
	).Scan(&id)
	if err != nil {
		return domain.Tensions{}, errors.Wrapf(err, "failed to insert tensions record")
	}
	tensions.Id = id
	return tensions, nil
}

func (s *Storage) CreateRecordSleep(sleep domain.Sleep) (domain.Sleep, error) {
	query := `
        INSERT INTO sleep (hours, minutes, control, date_of_addition, patient_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		sleep.Hours,
		sleep.Minutes,
		sleep.Control,
		sleep.DateOfAddition,
		sleep.IdPatient,
	).Scan(&id)
	if err != nil {
		return domain.Sleep{}, errors.Wrapf(err, "failed to insert sleep record")
	}
	sleep.Id = id
	return sleep, nil
}

func (s *Storage) CreateRecordWater(water domain.Water) (domain.Water, error) {
	query := `
        INSERT INTO water (volume_glass, count_glass, indicator, control, date_of_addition, patient_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		water.VolumeGlass,
		water.CountGlass,
		water.Indicator,
		water.Control,
		water.DateOfAddition,
		water.IdPatient,
	).Scan(&id)
	if err != nil {
		return domain.Water{}, errors.Wrapf(err, "failed to insert water record")
	}
	water.Id = id
	return water, nil
}

func (s *Storage) CreateRecordComplaints(complaints domain.Complaints) (domain.Complaints, error) {
	query := `
        INSERT INTO complaints (complaint, date_of_addition, patient_id)
        VALUES ($1, $2, $3)
        RETURNING id
    `
	var id int64
	err := s.connection.QueryRow(context.Background(), query,
		complaints.Complaint,
		complaints.DateOfAddition,
		complaints.IdPatient,
	).Scan(&id)
	if err != nil {
		return domain.Complaints{}, errors.Wrapf(err, "failed to insert complaints record")
	}
	complaints.Id = id
	return complaints, nil
}
