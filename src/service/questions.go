package service

import (
	"github.com/JrSchmidtt/covid-19-neo4J/src/storage"
)

// GetTotalCasesAndDeathsByCountryAndDate retrieves the total cases and deaths for a specific country and date.
func GetTotalCasesAndDeathsByCountryAndDate(country_code string, date string) ([]map[string]interface{}, error) {
	return storage.GetCovidByCountryAndDate(country_code, date)
}

// GetVaccinatedPeopleByCountryAndDate retrieves the total vaccinated people for a specific country and date.
func GetVaccinatedPeopleByCountryAndDate(country_code string, date string) (map[string]interface{}, error) {
	return storage.GetPersonsVaccinated1PlusDose(country_code, date)
}