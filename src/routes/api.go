package routes

import (
	"github.com/JrSchmidtt/covid-19-neo4J/src/handler"
	"github.com/JrSchmidtt/covid-19-neo4J/src/middleware"
	"github.com/gin-gonic/gin"
)

func API(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.Use(middleware.Log())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handler.HealthCheck)
	}
	etl := v1.Group("/etl") // Extract, Transform, Load
	{
		etl.POST("/csv", handler.ProcessCSV)
	}
	return r
}