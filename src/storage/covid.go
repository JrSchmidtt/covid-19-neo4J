package storage

import (
	"context"
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

func GetCovidByCountryAndDate(countryCode string, date string) ([]map[string]interface{}, error) {
	query := `
		MATCH (c:Covid)
		WHERE c.country_code = $code AND c.date_reported = $date_reported
		RETURN c
	`
	res, err := connection.ExecuteReadTransaction(context.Background(), query, map[string]interface{}{
		"code":          countryCode,
		"date_reported": date,
	})
	if err != nil {
		log.Println("Error getting covid record:", err)
		return nil, err
	}
	covid := res
	return covid, nil
}