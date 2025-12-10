package models

import "time"

type TripType string
type ComponentType string

const (
	Business TripType = "деловая поездка"
	Vacation TripType = "отпуск"

	Flight        ComponentType = "перелет"
	Accommodation ComponentType = "проживание"
	Excursion     ComponentType = "экскурсия"
	Transport     ComponentType = "транспорт"
)

type Component struct {
	ID          string        `json:"id"`
	Type        ComponentType `json:"type"`
	Cost        float64       `json:"cost"`
	StartDate   time.Time     `json:"start_date"`
	EndDate     time.Time     `json:"end_date"`
	Location    string        `json:"location"`
	ContactInfo string        `json:"contact_info"`
	Notes       string        `json:"notes"`
}

type Trip struct {
	ID          string     `json:"id"`
	Destination string     `json:"destination"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	Type        TripType   `json:"type"`
	Budget      float64    `json:"budget"`
	Components  []Component `json:"components"`
}
