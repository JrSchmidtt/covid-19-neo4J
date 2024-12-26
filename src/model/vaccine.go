package model

import "time"

// Represents a Vaccine along the time and world regions.
type Vaccine struct {
	ID                string    `json:"id"`
	ISO3              string    `json:"iso3" csv:"ISO3"`
	Product           string    `json:"product" csv:"PRODUCT_NAME"`
	Vaccine           string    `json:"vaccine" csv:"VACCINE_NAME"`
	Company           string    `json:"company" csv:"COMPANY_NAME"`
	AuthorizationDate string    `json:"authorization_date" csv:"AUTHORIZATION_DATE"`
	StartDate         string    `json:"start_date" csv:"START_DATE"`
	EndDate           string    `json:"end_date" csv:"END_DATE"`
	DataSource        string    `json:"data_source" csv:"DATA_SOURCE"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
