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

func CreateVaccine(vaccine model.Vaccine) (model.Vaccine, error) {
	vaccine.ID = id.NewIdWithPrefix("V")
	vaccine.CreatedAt = utils.TimeNow()
	vaccine.UpdatedAt = utils.TimeNow()

	vaccineMap, err := utils.StructToMap(vaccine, "json")
	if err != nil {
		log.Println("Error converting vaccine to map:", err)
		return model.Vaccine{}, err
	}

	query := `
    	MERGE (d:Date {date: $start_date})
    	WITH d

    	MERGE (c:Country {code: $country_code})
    	WITH d, c

		MERGE (v:Vaccine {product: $vaccine_product})
		ON CREATE SET v = $vaccine

    	MERGE (v)-[:STARTED_ON]->(d)
    	MERGE (v)-[:AUTHORIZATION_ON]->(authDate:Date {date: $authorization_date})
    	MERGE (c)-[:USES]->(v)
	`

	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"vaccine":            vaccineMap,
		"start_date":         vaccine.StartDate,
		"authorization_date": vaccine.AuthorizationDate,
		"country_code":       vaccine.ISO3,
		"vaccine_product":    vaccine.Product,
	})
	if err != nil {
		log.Println("Error creating vaccine record:", err)
		return model.Vaccine{}, err
	}

	return vaccine, nil
}

func GetVaccinesByCountryAndStartDate(country_code string, start_date string) (map[string]interface{}, error) {
	fmt	.Println("country_code", country_code)
	fmt.Println("start_date", start_date)
	query := `
		MATCH (v:Vaccine {iso3: $country_code})
		MATCH (v:Vaccine)-[:USES]->(c)
		MATCH (v)-[:STARTED_ON]->(d:Date {date: $start_date})
		RETURN c, v, d
	`

	result, err := connection.ExecuteReadTransactionMap(context.Background(), query, map[string]interface{}{
		"start_date": start_date,
		"country_code": country_code,
	})
	if err != nil {
		log.Println("Error retrieving vaccines by country and start date:", err)
		return nil, err
	}

	return result, nil
}


func GetMostUsedVaccineByRegion(region string) (map[string]interface{}, error) {
	query := `
		MATCH (r:Region {name: $region})
		MATCH (c:Country)-[:BELONGS]->(r)
		MATCH (c)-[:USES]->(v:Vaccine)
		RETURN v.vaccine AS vaccine, COUNT(v) AS count
		ORDER BY count DESC
		LIMIT 1
	`

	result, err := connection.ExecuteReadTransactionMap(context.Background(), query, map[string]interface{}{
		"region": region,
	})
	if err != nil {
		log.Println("Error retrieving most used vaccine by region:", err)
		return nil, err
	}
	return result, nil
}