package storage

import (
	"context"
	"log"

	"github.com/JrSchmidtt/covid-19-neo4J/pkg/id"
	"github.com/JrSchmidtt/covid-19-neo4J/src/connection"
	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
)


func CreateCountry(country model.Country) (model.Country, error) {
	country.ID = id.NewIdWithPrefix("C")
	country.CreatedAt = utils.TimeNow()
	country.UpdatedAt = utils.TimeNow()

	countryMap, err := utils.StructToMap(country, "json")
	if err != nil {
		log.Println("Error converting country to map:", err)
		return model.Country{}, err
	}

	query := `
		MERGE (r:Region {name: $region})

		CREATE (c:Country)
		SET c = $country

		MERGE (c)-[:BELONGS]->(r)
	`

	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"country": countryMap,
		"region": country.Region,
	})
	if err != nil {
		log.Println("Error creating country record:", err)
		return model.Country{}, err
	}
	return country, nil
}


func GetCountryWithHighestCasesByDate(date string) (map[string]interface{}, error) {
	query := `
		MATCH (c:Country)-[:HAS]->(covid:Covid)
		WHERE covid.date = $date
		WITH c, covid
		ORDER BY covid.total_cases DESC
		LIMIT 1
		RETURN c.name AS country, covid.total_cases AS total_cases
	`

	result, err := connection.ExecuteReadTransactionMap(context.Background(), query, map[string]interface{}{
		"date": date,
	})
	if err != nil {
		log.Println("Error getting country with highest cases by date:", err)
		return nil, err
	}
	return result, nil
}