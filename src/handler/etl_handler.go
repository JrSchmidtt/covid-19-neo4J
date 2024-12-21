package handler

import (
	"net/http"
	"github.com/JrSchmidtt/covid-19-neo4J/src/service"
	"github.com/gin-gonic/gin"
)

var comma = ';'
var commaCovidGlobal = ','

func ProcessCSV(c *gin.Context) {
	vaccinationFile, _, err := c.Request.FormFile("vaccination");if err != nil {
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'vaccination'"})
		return
	}
	defer vaccinationFile.Close()
	vaccinationRaw, err := service.ProcessCSV(vaccinationFile, comma); if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	vaccineFile, _, err := c.Request.FormFile("vaccine"); if err != nil {
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'vaccine'"})
		return
	}
	defer vaccineFile.Close()
	vaccineRaw, err := service.ProcessCSV(vaccineFile, comma); if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	covidFile, _, err := c.Request.FormFile("covid"); if err != nil {
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'covid'"})
		return
	}
	defer covidFile.Close()
	covidRaw, err := service.ProcessCSV(covidFile, comma); if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	covidGlobalFile, _, err := c.Request.FormFile("covid_global")
	if err != nil {
		c.JSON(400, gin.H{"error": "bad request please provide a file with key 'covid_global'"})
		return
	}
	defer covidGlobalFile.Close()
	covidGlobaRaw, err := service.ProcessCSV(covidGlobalFile, commaCovidGlobal)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"vaccinationRaw": vaccinationRaw,
		"vaccineRaw":     vaccineRaw,
		"covidRaw":       covidRaw,
		"covidGlobalRaw": covidGlobaRaw,
	})
}
