package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	once       sync.Once
	Neo4jDriver *neo4j.DriverWithContext
	Neo4Ctx   context.Context
	cancelFunc  context.CancelFunc
)

// InitNeo4j initializes the Neo4j driver and global context as a Singleton
func InitNeo4j() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file but don't worry, we are using the environment variables.")
		}

		uri := os.Getenv("NEO4J_URI")
		username := os.Getenv("NEO4J_USER")
		password := os.Getenv("NEO4J_PASSWORD")

		// Singleton context with a long timeout (e.g., for the lifetime of the application)
		Neo4Ctx, cancelFunc = context.WithTimeout(context.Background(), 60*time.Minute)

		driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
		if err != nil {
			log.Fatal("Error creating Neo4j driver: ", err)
		}

		// Test the connection
		if err = testConnection(Neo4Ctx, &driver); err != nil {
			log.Fatal("Error connecting to Neo4j: ", err)
		}

		Neo4jDriver = &driver
		fmt.Println("Connected to Neo4j!")
	})
}

// testConnection pings the Neo4j database to validate the connection
func testConnection(ctx context.Context, driver *neo4j.DriverWithContext) error {
	session := (*driver).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		return tx.Run(ctx, "RETURN 1", nil)
	})
	return err
}

// GetNeo4jDriver returns the Singleton Neo4j driver instance
func GetNeo4jDriver() *neo4j.DriverWithContext {
	if Neo4jDriver == nil {
		InitNeo4j()
	}
	return Neo4jDriver
}

// GetGlobalContext returns the Singleton context
func GetNeo4jContext() context.Context {
	if Neo4Ctx == nil {
		InitNeo4j()
	}
	return Neo4Ctx
}

// CloseNeo4j closes the Neo4j driver and cancels the global context
func CloseNeo4j() {
	if Neo4jDriver != nil {
		(*Neo4jDriver).Close(Neo4Ctx)
	}
	if cancelFunc != nil {
		cancelFunc()
	}
	fmt.Println("Neo4j connection closed.")
}
