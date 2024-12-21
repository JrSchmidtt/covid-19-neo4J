package handler

import (
	"encoding/csv"
	"log"
	"net/http"

	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
	"github.com/gin-gonic/gin"
)

var comma = ';' // default comma separator
var commaCovidGlobal = ','

func ProcessCSV(c *gin.Context) {
	vaccinationDataFile, _, err := c.Request.FormFile("vaccination")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'vaccination'"})
		return
	}
	defer vaccinationDataFile.Close()
	vaccinationReader := csv.NewReader(vaccinationDataFile)
	vaccinationReader.Comma = comma
	vaccinationRawData, err := utils.ExtractDataFromCSV(vaccinationReader)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	vaccineDataFile, _, err := c.Request.FormFile("vaccine")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'vaccine'"})
		return
	}
	defer vaccineDataFile.Close()
	vaccineReader := csv.NewReader(vaccineDataFile)
	vaccineReader.Comma = comma
	vaccineRawData, err := utils.ExtractDataFromCSV(vaccineReader)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	covidDataFile, _, err := c.Request.FormFile("covid")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'covid'"})
		return
	}
	defer covidDataFile.Close()
	covidReader := csv.NewReader(covidDataFile)
	covidReader.Comma = comma
	covidRawData, err := utils.ExtractDataFromCSV(covidReader)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	covidGlobalDataFile, _, err := c.Request.FormFile("covid_global")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'covid_global'"})
		return
	}
	defer covidGlobalDataFile.Close()
	covidGlobalReader := csv.NewReader(covidGlobalDataFile)
	covidGlobalReader.Comma = commaCovidGlobal
	covidGlobalRawData, err := utils.ExtractDataFromCSV(covidGlobalReader)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vaccinationRawData": vaccinationRawData,
		"vaccineRawData":     vaccineRawData,
		"covidRawData":       covidRawData,
		"covidGlobalRawData": covidGlobalRawData,
	})
}
