package connection

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"os"
	"sync"
	"time"
)

var (
	once         sync.Once
	Neo4jDriver  *neo4j.DriverWithContext
	sessionOnce  sync.Once
	Neo4jSession *neo4j.SessionWithContext
	mu           sync.Mutex
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
	if Neo4jDriver == nil {
		InitNeo4j()
	}
	return Neo4jDriver
}

// GetNeo4jSession returns a Singleton instance of the Neo4j session
func GetNeo4jSession() *neo4j.SessionWithContext {
	sessionOnce.Do(func() {
		driver := GetNeo4jDriver()
		session := (*driver).NewSession(context.Background(), neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
		Neo4jSession = &session
	})
	return Neo4jSession
}

// CloseNeo4j closes the Neo4j driver and the session
func CloseNeo4j() {
	mu.Lock()
	defer mu.Unlock()

	if Neo4jSession != nil {
		(*Neo4jSession).Close(context.Background())
		Neo4jSession = nil
		fmt.Println("Neo4j session closed.")
	}

	if Neo4jDriver != nil {
		(*Neo4jDriver).Close(context.Background())
		Neo4jDriver = nil
		fmt.Println("Neo4j connection closed.")
	}
}

// ExecuteWriteTransaction executes a write transaction in Neo4j using the Singleton session
func ExecuteWriteTransaction(ctx context.Context, query string, params map[string]interface{}) (any, error) {
	session := GetNeo4jSession()

	data, err := (*session).ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})
	if err != nil {
		if neo4j.IsNeo4jError(err) && err.(*neo4j.Neo4jError).Code == "Neo.ClientError.Schema.ConstraintValidationFailed" {

			return nil, nil 
		}
		log.Println("Error executing write transaction:", err)
		return nil, err
	}
	return data, nil
}


// ExecuteReadTransaction executes a read transaction in Neo4j using the Singleton session
func ExecuteReadTransaction(ctx context.Context, query string, params map[string]interface{}) ([]map[string]interface{}, error) {
	session := GetNeo4jSession()
	var nodes []map[string]interface{}
	_, err := (*session).ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		for result.Next(ctx) {
			record := result.Record()
			node, _ := record.Get("c")
			if node != nil {
				neoNode := node.(neo4j.Node)
				properties := neoNode.Props
				nodes = append(nodes, properties)
			}
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		fmt.Println("Nodes:", nodes)
		return nodes, nil
	})
	if err != nil {
		log.Println("Error executing read transaction:", err)
		return nil, err
	}
	return nodes, nil
}

// ExecuteReadTransaction executes a read transaction in Neo4j using the Singleton session
func ExecuteReadTransactionMap(ctx context.Context, query string, params map[string]interface{}) (map[string]interface{}, error) {
	session := GetNeo4jSession()
	resultMap := make(map[string]interface{})
	_, err := (*session).ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		for result.Next(ctx) {
			record := result.Record().AsMap()
			for key, value := range record {
				resultMap[key] = value
			}
		}
		if result.Err() != nil {
			return nil, result.Err()
		}
		return result, nil
	})

	if err != nil {
		log.Println("Error executing read transaction:", err)
		return nil, err
	}
	if len(resultMap) == 0 {
		return nil, fmt.Errorf("no records found")
	}
	return resultMap, nil
}