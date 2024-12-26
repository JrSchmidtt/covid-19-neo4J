package storage

import (
	"context"
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

		CREATE (v:Vaccine)
		SET v = $vaccine

		MERGE (v)-[:STARTED_ON]->(d)
		MERGE (v)-[:AUTHORIZATION_ON]->(authDate:Date {date: $authorization_date})
	`

	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"vaccine":            vaccineMap,
		"start_date":         vaccine.StartDate,
		"authorization_date": vaccine.AuthorizationDate,
	})
	if err != nil {
		log.Println("Error creating vaccine record:", err)
		return model.Vaccine{}, err
	}

	return vaccine, nil
}

