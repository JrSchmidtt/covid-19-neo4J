package model

import "time"

// Represents a Covid-19 along the time and world regions.
type Covid struct {
	ID               string    `json:"id"` // Internal ID for the record
	DateReported     string    `json:"date_reported" csv:"Date_reported"`
	CountryCode      string    `json:"country_code" csv:"Country_code"`
	Country          string    `json:"country" csv:"Country"`
	WHORegion        string    `json:"who_region" csv:"WHO_region"`
	NewCases         int64     `json:"new_cases" csv:"New_cases"`
	CumulativeCases  int64     `json:"cumulative_cases" csv:"Cumulative_cases"`
	NewDeaths        int64     `json:"new_deaths" csv:"New_deaths"`
	CumulativeDeaths int64     `json:"cumulative_deaths" csv:"Cumulative_deaths"`
	CreatedAt        time.Time `json:"created_at"` // Date when the record was created
	UpdatedAt        time.Time `json:"updated_at"` // Date when the record was updated
}

// Represents a Covid-19 global data.
type CovidGlobal struct {
	ID                      string    `json:"id"` // Internal ID for the record
	Country                 string    `json:"country" csv:"Name"`
	WHORegion               string    `json:"who_region" csv:"WHO Region"`
	CumulativeCases         int64     `json:"cumulative_cases" csv:"Cases - cumulative total"`
	CumulativeCasesPer100k  int64     `json:"cumulative_cases_per100k" csv:"Cases - cumulative total per 100000 population"`
	NewCases7Days           int64     `json:"new_cases_7_days" csv:"Cases - newly reported in last 7 days"`
	NewCases7DaysPer100k    int64     `json:"new_cases_7_days_per100k" csv:"Cases - newly reported in last 7 days per 100000 population"`
	NewCases24Hours         int64     `json:"new_cases_24_hours" csv:"Cases - newly reported in last 24 hours"`
	CumulativeDeaths        int64     `json:"cumulative_deaths" csv:"Deaths - cumulative total"`
	CumulativeDeathsPer100k int64     `json:"cumulative_deaths_per100k" csv:"Deaths - cumulative total per 100000 population"`
	NewDeaths7Days          int64     `json:"new_deaths_7_days" csv:"Deaths - newly reported in last 7 days"`
	NewDeaths7DaysPer100k   int64     `json:"new_deaths_7_days_per100k" csv:"Deaths - newly reported in last 7 days per 100000 population"`
	NewDeaths24Hours        int64     `json:"new_deaths_24_hours" csv:"Deaths - newly reported in last 24 hours"`
	CreatedAt               time.Time `json:"created_at"` // Date when the record was created
	UpdatedAt               time.Time `json:"updated_at"` // Date when the record was updated
}
