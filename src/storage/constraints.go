package storage

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/JrSchmidtt/covid-19-neo4J/src/connection"
)

// CreateDatabaseConstraints creates the necessary constraints in the Neo4j database
func CreateDatabaseConstraints() error {
	constraints := []string{
		// Node Keys
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Vaccine) REQUIRE n.product IS NODE KEY`,
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Date) REQUIRE n.date IS NODE KEY`,
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Region) REQUIRE n.name IS NODE KEY`,
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Country) REQUIRE n.code IS NODE KEY`,

		// Node Properties Must Exist
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Vaccine) REQUIRE n.company IS NOT NULL`,
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Vaccine) REQUIRE n.vaccine IS NOT NULL`,
		`CREATE CONSTRAINT IF NOT EXISTS FOR (n:Country) REQUIRE n.name IS NOT NULL`,
	}
	for _, constraint := range constraints {
		_, err := connection.ExecuteWriteTransaction(context.Background(), constraint, nil)
		if err != nil {
			if strings.Contains(err.Error(), "Constraint already exists") {
				log.Printf("Constraint already exists, skipping creation: %v\n", err)
			} else {
				log.Printf("Error creating constraint: %v\n", err)
				return err
			}
		} else {
			fmt.Println("Constraint created successfully.")
		}
	}

	return nil
}