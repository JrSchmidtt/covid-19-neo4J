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
	questions := v1.Group("/questions")
	{
		// 1. Total acumulado de casos e mortes de Covid-19 em um país específico em uma data determinada
		questions.GET("/1", handler.GetTotalCasesAndDeathsByCountryAndDate)

		// 2. Pessoas vacinadas com pelo menos uma dose em um país específico em uma data específica
		questions.GET("/2", handler.GetVaccinatedPeopleByCountryAndDate)

		// // 3. Vacinas usadas em um país e as datas de início de aplicação
		// questions.GET("/3", handler.GetVaccinesByCountryAndStartDate)

		// // 4. País com o maior número de casos acumulados até uma data específica
		// questions.GET("/4", handler.GetCountryWithHighestCasesByDate)

		// // 5. Vacina mais utilizada em uma região específica
		// questions.GET("/5", handler.GetMostUsedVaccineByRegion)
	}
	return r
}
