package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/JrSchmidtt/covid-19-neo4J/pkg/id"
	"github.com/JrSchmidtt/covid-19-neo4J/src/connection"
	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
)

// CreateVaccination creates a new vaccination record and relates it to a specific date by first_vaccine_date.
func CreateVaccination(vaccination model.Vaccination) (model.Vaccination, error) {
	vaccination.CreatedAt = utils.TimeNow()
	vaccination.UpdatedAt = utils.TimeNow()
	vaccination.ID = id.NewIdWithPrefix("V")

	vaccinationMap, err := utils.StructToMap(vaccination, "json")
	if err != nil {
		log.Println("Error converting vaccination to map:", err)
		return model.Vaccination{}, err
	}

	query := `
		MERGE (d:Date {date: $first_vaccine_date})

		WITH d

		MATCH (c:Country {name: $country})

		CREATE (v:Vaccination)
		SET v = $vaccination
		
		MERGE (v)-[:ON_DATE]->(d)
		MERGE (v)-[:VACCINATED_ON]->(c)
	`

	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"vaccination":        vaccinationMap,
		"first_vaccine_date": vaccination.FirstVaccineDate,
		"country":            vaccination.Country,
	})
	if err != nil {
		log.Println("Error creating vaccination record:", err)
		return model.Vaccination{}, err
	}
	return vaccination, nil
}

// GetPersonsVaccinated1PlusDose retrieves the number of people vaccinated with at least one dose for a specific country and date.
func GetPersonsVaccinated1PlusDose(countryCode string, date string) (map[string]interface{}, error) {
	query := `
		MATCH (c:Country {code: $country_code})
		MATCH (v:Vaccination)-[:VACCINATED_ON]->(c)
		MATCH (v)-[:ON_DATE]->(d:Date {date: $date})
		RETURN c, v, d
	`
	result, err := connection.ExecuteReadTransactionMap(context.Background(), query, map[string]interface{}{
		"country_code": countryCode,
		"date":         date,
	})
	if err != nil {
		log.Println("Error getting persons vaccinated 1+ dose:", err)
		return nil, err
	}
	fmt.Println("Result:", result)
	return result, nil
}
