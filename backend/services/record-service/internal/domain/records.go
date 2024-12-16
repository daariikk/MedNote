package domain

type InfoField struct {
	Id             int64  `json:"id"`
	Control        string `json:"control"`
	DateOfAddition string `json:"date_of_addition"`
	IdPatient      int64  `json:"patient_id"`
}

type Pulse struct {
	Indicator int64 `json:"indicator"`
	InfoField
}

type Steps struct {
	Indicator int64 `json:"indicator"`
	InfoField
}

type Tensions struct {
	UpperIndicator int64 `json:"upper_indicator"`
	LowerIndicator int64 `json:"lower_indicator"`
	InfoField
}

type Sleep struct {
	Hours   int64 `json:"hours"`
	Minutes int64 `json:"minutes"`
	InfoField
}

type Water struct {
	VolumeGlass int64 `json:"volume_glass"`
	CountGlass  int64 `json:"count_glass"`
	Indicator   int64 `json:"indicator"`
	InfoField
}

type Complaints struct {
	Id             int64  `json:"id"`
	Complaint      string `json:"complaint"`
	DateOfAddition string `json:"date_of_addition"`
	IdPatient      int64  `json:"patient_id"`
}
