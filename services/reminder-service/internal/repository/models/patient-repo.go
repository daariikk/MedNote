package models

import (
	"github.com/daariikk/MedNote/services/patient-service/internal/domain"
	"time"
)

// Не успела сделать разницу уровней моделек для бд
type Patient struct {
	Id               int64     `json:"patient_id" required:"true"`
	FirstName        string    `json:"first_name"`
	SecondName       string    `json:"second_name"`
	Email            string    `json:"email" unique:"true"`
	Height           float64   `json:"height"`
	Weight           float64   `json:"weight"`
	Gender           string    `json:"gender"`
	Password         []byte    `json:"password"`
	RegistrationData time.Time `json:"registration_data"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// TODO: храню на стороне клиента в jwt токене claims id
// TODO: храню пароли в базе в виде Hash HSA256 в формате base64
func DomainToRepo(patient domain.Patient) Patient {
	return Patient{
		Id:               patient.Id,
		FirstName:        patient.FirstName,
		SecondName:       patient.SecondName,
		Email:            patient.Email,
		Height:           patient.Height,
		Weight:           patient.Weight,
		Gender:           patient.Gender,
		Password:         patient.Password,
		RegistrationData: patient.RegistrationData,
	}
}

func RepoToDomain(patient Patient) domain.Patient {
	return domain.Patient{
		Id:               patient.Id,
		FirstName:        patient.FirstName,
		SecondName:       patient.SecondName,
		Email:            patient.Email,
		Height:           patient.Height,
		Weight:           patient.Weight,
		Gender:           patient.Gender,
		Password:         patient.Password,
		RegistrationData: patient.RegistrationData,
	}
}
