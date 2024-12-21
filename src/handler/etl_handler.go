package handler

import (
	"encoding/csv"
	"log"
	"net/http"

	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
	"github.com/gin-gonic/gin"
)

var comma = ';' // default comma separator

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
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'vaccine-data'"})
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
	c.JSON(http.StatusOK, gin.H{
		"vaccinationRawData": vaccinationRawData,
		"vaccineRawData":     vaccineRawData,
	})
}
