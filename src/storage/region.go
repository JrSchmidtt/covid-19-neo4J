package storage

import (
	"context"
	"log"

	"github.com/JrSchmidtt/covid-19-neo4J/pkg/id"
	"github.com/JrSchmidtt/covid-19-neo4J/src/connection"
	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
)

func CreateRegion(region model.Region) (model.Region, error) {
	region.ID = id.NewIdWithPrefix("R")
	region.CreatedAt = utils.TimeNow()
	region.UpdatedAt = utils.TimeNow()

	regionMap, err := utils.StructToMap(region, "json")
	if err != nil {
		log.Println("Error converting region to map:", err)
		return model.Region{}, err
	}

	query := `
		CREATE (r:Region)
		SET r = $region
	`

	_, err = connection.ExecuteWriteTransaction(context.Background(), query, map[string]interface{}{
		"region": regionMap,
	})
	if err != nil {
		log.Println("Error creating region record:", err)
		return model.Region{}, err
	}
	return region, nil
}