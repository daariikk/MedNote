package domain

import "time"

type User struct {
	Id               int64     `json:"patient_id"`
	FirstName        string    `json:"first_name"`
	SecondName       string    `json:"second_name,omitempty"`
	Email            string    `json:"email"`
	Height           float64   `json:"height,omitempty"`
	Weight           float64   `json:"weight,omitempty"`
	Gender           string    `json:"gender,omitempty"`
	Password         string    `json:"password,omitempty"`
	RegistrationData time.Time `json:"registration_data,omitempty"`
}
