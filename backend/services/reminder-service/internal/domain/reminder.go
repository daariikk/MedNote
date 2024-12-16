package domain

type Reminder struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	PatientId int64  `json:"patient_id"`
}
