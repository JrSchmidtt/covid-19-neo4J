package model

import "time"

// Represents a Vaccine along the time and world regions.
type Vaccine struct {
	ID                string    `json:"id"`
	ISO3              string    `json:"iso3"`
	Product           string    `json:"product"`
	Vaccine           string    `json:"vaccine"`
	Company           string    `json:"company"`
	AuthorizationDate time.Time `json:"authorization_date"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	DataSource        string    `json:"data_source"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
