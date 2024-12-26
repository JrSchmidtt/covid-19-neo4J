package server

import (
	"fmt"
	"log"
	"os"

	"github.com/JrSchmidtt/covid-19-neo4J/src/connection"
	"github.com/JrSchmidtt/covid-19-neo4J/src/routes"
	"github.com/JrSchmidtt/covid-19-neo4J/src/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() error {
	err := godotenv.Load() ; if err != nil {
		log.Fatal(err)
	}
	connection.InitNeo4j(); if err != nil {
		log.Fatal(err)
	}
	storage.CreateDatabaseConstraints(); if err != nil {
		log.Fatal(err)
	}
	gin := gin.Default()
	router := routes.API(gin)
	err = router.Run(":" + os.Getenv("PORT"))
	fmt.Println("Server running on port", os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
