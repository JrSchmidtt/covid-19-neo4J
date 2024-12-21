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
	once        sync.Once
	Neo4jDriver *neo4j.DriverWithContext
	mu          sync.Mutex
)

// InitNeo4j initializes the Neo4j driver as a Singleton
func InitNeo4j() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file but don't worry, we are using the environment variables.")
		}

		uri := os.Getenv("NEO4J_URI")
		username := os.Getenv("NEO4J_USER")
		password := os.Getenv("NEO4J_PASSWORD")

		driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
		if err != nil {
			log.Fatal("Error creating Neo4j driver: ", err)
		}

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = testConnection(ctx, &driver); err != nil {
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

// GetOperationContext creates a new context with a short timeout for a single operation
func GetOperationContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// GetNeo4jDriver returns the Singleton Neo4j driver instance
func GetNeo4jDriver() *neo4j.DriverWithContext {
	mu.Lock()
	defer mu.Unlock()

	if Neo4jDriver == nil {
		InitNeo4j()
	}
	return Neo4jDriver
}

// CloseNeo4j closes the Neo4j driver
func CloseNeo4j() {
	mu.Lock()
	defer mu.Unlock()

	if Neo4jDriver != nil {
		(*Neo4jDriver).Close(context.Background())
		fmt.Println("Neo4j connection closed.")
	}
}
