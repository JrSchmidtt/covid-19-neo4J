package utils

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// MapNeo4jNodeToStruct maps a Neo4j node from a result to a struct.
// res is the Neo4j result map
// nodeKey is the key of the node to be mapped
// result is the struct to be mapped
func Neo4jNodeToStruct(res map[string]interface{}, nodeKey string, result interface{}) error {
	node, ok := res[nodeKey].(neo4j.Node)
	if !ok {
		return fmt.Errorf("error asserting %s data, expected a map[string]interface{}", nodeKey)
	}
	return MapToStruct(node.Props, result)
}