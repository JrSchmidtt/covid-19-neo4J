package handler

import (
	"net/http"

	"github.com/JrSchmidtt/covid-19-neo4J/src/service"
	"github.com/gin-gonic/gin"
)

func GetTotalCasesAndDeathsByCountryAndDate(c *gin.Context) {
	country_code := c.Query("country_code")
	date := c.Query("date")
	result, err := service.GetTotalCasesAndDeathsByCountryAndDate(country_code, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": result,
	})
}

func GetVaccinatedPeopleByCountryAndDate(c *gin.Context) {
	country_code := c.Query("country_code")
	date := c.Query("date")
	result, err := service.GetVaccinatedPeopleByCountryAndDate(country_code, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": result,
	})
}