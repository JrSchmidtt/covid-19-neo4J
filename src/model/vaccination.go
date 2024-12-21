package model

import "time"

// Represents a vaccination data from a country
type VaccinationStats struct {
	ID                               string    `json:"id"` // Internal ID for the record
	Country                          string    `json:"country"`
	ISO3                             string    `json:"iso3"`
	WHORegion                        string    `json:"who_region"`  // World Health Organization Region
	DataSource                       string    `json:"data_source"` // WPRO, REPORTING, EMRO, etc
	TotalVaccinations                int64     `json:"total_vaccinations"`
	PersonsVaccinated1PlusDose       int64     `json:"persons_vaccinated_1plus_dose"`
	TotalVaccinationsPer100          int64     `json:"total_vaccinations_per100"`
	PersonsVaccinated1PlusDosePer100 int64     `json:"persons_vaccinated_1plus_dose_per100"`
	PersonsLastDose                  int64     `json:"persons_last_dose"`
	PersonsLastDosePer100            int64     `json:"persons_last_dose_per100"`
	NumberVaccinesTypesUsed          int64     `json:"number_vaccines_types_used"`
	PersonsBoosterAddDose            int64     `json:"persons_booster_add_dose"`
	PersonsBoosterAddDosePer100      int64     `json:"persons_booster_add_dose_per100"`
	DateUpdated                      time.Time `json:"date_updated"`
	FirstVaccineDate                 time.Time `json:"first_vaccine_date"`
	CreatedAt                        time.Time `json:"created_at"` // Date when the record was created
	UpdatedAt                        time.Time `json:"updated_at"` // Date when the record was updated
}