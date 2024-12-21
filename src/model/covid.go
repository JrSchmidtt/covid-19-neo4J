package model

import "time"

// Represents a Covid-19 along the time and world regions.
type CovidStats struct {
	ID               string    `json:"id"` // Internal ID for the record
	DateReported     time.Time `json:"date_reported"`
	CountryCode      string    `json:"country_code"`
	Country          string    `json:"country"`
	WHORegion        string    `json:"who_region"`
	NewCases         int64     `json:"new_cases"`
	CumulativeCases  int64     `json:"cumulative_cases"`
	NewDeaths        int64     `json:"new_deaths"`
	CumulativeDeaths int64     `json:"cumulative_deaths"`
	CreatedAt        time.Time `json:"created_at"` // Date when the record was created
	UpdatedAt        time.Time `json:"updated_at"` // Date when the record was updated
}