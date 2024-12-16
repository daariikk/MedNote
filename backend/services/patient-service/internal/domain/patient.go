package domain

import "time"

// validation for this model is not required - frontend task
type Patient struct {
	Id               int64     `json:"patient_id"`
	FirstName        string    `json:"first_name"`
	SecondName       string    `json:"second_name"`
	Email            string    `json:"email"`
	Height           float64   `json:"height"`
	Weight           float64   `json:"weight"`
	Gender           string    `json:"gender"`
	Password         string    `json:"password"`
	RegistrationData time.Time `json:"registration_data"`
}
