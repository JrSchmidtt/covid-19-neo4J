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

func CreateCovid(covid model.Covid) (model.Covid, error) {
	covid.ID = id.NewIdWithPrefix("C")
	covid.CreatedAt = utils.TimeNow()
	covid.UpdatedAt = utils.TimeNow()

	covidMap, err := utils.StructToMap(covid, "json")
	if err != nil {
		log.Println("Error converting covid to map:", err)
		return model.Covid{}, err
	}

	query := `
		MERGE (n:Country {name: $country})
		MERGE (d:Date {date: $date_reported})

		CREATE (c:Covid)
		SET c = $covid

		MERGE (c)-[:REPORTED_ON]->(n)
		MERGE (c)-[:ON_DATE]->(d)
	`
	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"covid":         covidMap,
		"country":       covid.Country,
		"date_reported": covid.DateReported,
	})
	if err != nil {
		log.Println("Error creating covid record:", err)
		return model.Covid{}, err
	}
	return covid, nil
}

func CreateCovidGlobal(covidGlobal model.CovidGlobal) (model.CovidGlobal, error) {
	covidGlobal.ID = id.NewIdWithPrefix("CG")
	covidGlobal.CreatedAt = utils.TimeNow()
	covidGlobal.UpdatedAt = utils.TimeNow()

	covidGlobalMap, err := utils.StructToMap(covidGlobal, "json")
	if err != nil {
		log.Println("Error converting covid global to map:", err)
		return model.CovidGlobal{}, err
	}

	query := `
		MERGE (c:Country {name: $country})

		CREATE (cg:CovidGlobal)
		SET cg = $covid_global

		MERGE (cg)-[:REPORTED_ON]->(c)
	`
	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"covid_global": covidGlobalMap,
		"country":      covidGlobal.Country,
	})
	if err != nil {
		log.Println("Error creating covid global record:", err)
		return model.CovidGlobal{}, err
	}
	return covidGlobal, nil
}

func GetCovidByCountryAndDate(countryCode string, date string) (model.Covid, error) {
	query := `
		MATCH (d:Date {date: $date})
		MATCH (cv:Covid)-[:ON_DATE]->(d)
		MATCH (cv:Covid)-[:REPORTED_ON]->(c:Country {code: $code})
		RETURN cv
	`
	res, err := connection.ExecuteReadTransactionMap(context.Background(), query, map[string]interface{}{
		"code": countryCode,
		"date": date,
	})
	if err != nil {
		return model.Covid{}, fmt.Errorf("error getting covid record: %v", err)
	}
	var covid model.Covid
	err = utils.Neo4jNodeToStruct(res, "cv", &covid); if err != nil {
		return model.Covid{}, fmt.Errorf("error mapping country node to struct: %v", err)
	}
	return covid, nil
}