package handler

import (
	"fmt"
	"net/http"

	"github.com/JrSchmidtt/covid-19-neo4J/src/service"
	"github.com/gin-gonic/gin"
)

func GetTotalCasesAndDeathsByCountryAndDate(c *gin.Context) {
	country_code := c.Query("country_code")
	date := c.Query("date")
	covid, err := service.GetTotalCasesAndDeathsByCountryAndDate(country_code, date)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"covid": covid,
	})
}

func GetVaccinatedPeopleByCountryAndDate(c *gin.Context) {
	country_code := c.Query("country_code")
	date := c.Query("date")
	result, err := service.GetVaccinatedPeopleByCountryAndDate(country_code, date)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": result,
	})
}

func GetVaccinesByCountryAndStartDate(c *gin.Context) {
	country_code := c.Query("country_code")
	start_date := c.Query("start_date")
	fmt.Println("country_code", country_code)
	fmt.Println("start_date", start_date)
	result, err := service.GetVaccinesByCountryAndStartDate(country_code, start_date)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": result,
	})
}

func GetMostUsedVaccineByRegion(c *gin.Context) {
	region := c.Query("region")
	result, err := service.GetMostUsedVaccineByRegion(region)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": result,
	})
}

func GetCountryWithHighestCasesByDate(c *gin.Context) {
	date := c.Query("date")
	result, err := service.GetCountryWithHighestCasesByDate(date)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": result,
	})
}