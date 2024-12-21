package model

import "time"

// Represents a Covid-19 along the time and world regions.
type Covid struct {
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

// Represents a Covid-19 global data.
type CovidGlobal struct {
	ID                      string    `json:"id"` // Internal ID for the record
	Country                 string    `json:"country"`
	WHORegion               string    `json:"who_region"`
	CumulativeCases         int64     `json:"cumulative_cases"`
	CumulativeCasesPer100k  int64     `json:"cumulative_cases_per100k"`
	NewCases7Days           int64     `json:"new_cases_7_days"`
	NewCases7DaysPer100k    int64     `json:"new_cases_7_days_per100k"`
	NewCases24Hours         int64     `json:"new_cases_24_hours"`
	CumulativeDeaths        int64     `json:"cumulative_deaths"`
	CumulativeDeathsPer100k int64     `json:"cumulative_deaths_per100k"`
	NewDeaths7Days          int64     `json:"new_deaths_7_days"`
	NewDeaths7DaysPer100k   int64     `json:"new_deaths_7_days_per100k"`
	NewDeaths24Hours        int64     `json:"new_deaths_24_hours"`
	CreatedAt               time.Time `json:"created_at"` // Date when the record was created
	UpdatedAt               time.Time `json:"updated_at"` // Date when the record was updated
}