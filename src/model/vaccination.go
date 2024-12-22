package model

import "time"

// Represents a vaccination data from a country or world region.
type Vaccination struct {
	ID                               string    `json:"id" csv:"ID"` // Internal ID for the record
	Country                          string    `json:"country" csv:"COUNTRY"`
	ISO3                             string    `json:"iso3" csv:"ISO3"`
	WHORegion                        string    `json:"who_region" csv:"WHO_REGION"`   // World Health Organization Region
	DataSource                       string    `json:"data_source" csv:"DATA_SOURCE"` // WPRO, REPORTING, EMRO, etc
	TotalVaccinations                float64   `json:"total_vaccinations" csv:"TOTAL_VACCINATIONS"`
	PersonsVaccinated1PlusDose       float64   `json:"persons_vaccinated_1plus_dose" csv:"PERSONS_VACCINATED_1PLUS_DOSE"`
	TotalVaccinationsPer100          int64     `json:"total_vaccinations_per100" csv:"TOTAL_VACCINATIONS_PER100"`
	PersonsVaccinated1PlusDosePer100 int64     `json:"persons_vaccinated_1plus_dose_per100" csv:"PERSONS_VACCINATED_1PLUS_DOSE_PER100"`
	PersonsLastDose                  int64     `json:"persons_last_dose" csv:"PERSONS_LAST_DOSE"`
	PersonsLastDosePer100            int64     `json:"persons_last_dose_per100" csv:"PERSONS_LAST_DOSE_PER100"`
	PersonsBoosterAddDose            int64     `json:"persons_booster_add_dose" csv:"PERSONS_BOOSTER_ADD_DOSE"`
	PersonsBoosterAddDosePer100      int64     `json:"persons_booster_add_dose_per100" csv:"PERSONS_BOOSTER_ADD_DOSE_PER100"`
	// DateUpdated                      time.Time `json:"date_updated" csv:"DATE_UPDATED,parseTime" time_format:"02/01/2006"`
	// FirstVaccineDate                 time.Time `json:"first_vaccine_date" csv:"FIRST_VACCINE_DATE,parseTime" time_format:"02/01/2006"`
	CreatedAt                        time.Time `json:"created_at" csv:"CREATED_AT"` // Date when the record was created
	UpdatedAt                        time.Time `json:"updated_at" csv:"UPDATED_AT"` // Date when the record was updated
}
