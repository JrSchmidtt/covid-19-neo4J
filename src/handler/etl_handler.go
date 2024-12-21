package handler

import (
	"net/http"
	"github.com/JrSchmidtt/covid-19-neo4J/src/service"
	"github.com/gin-gonic/gin"
)

var comma = ';'
var commaCovidGlobal = ','

func ProcessCSV(c *gin.Context) {
	var errors []string
	var warnings []string

	vaccinationRaw, err := processFile(c, "vaccination", comma)
	if err != nil {
		errors = append(errors, "Error processing 'vaccination': "+err.Error())
	}

	vaccineRaw, err := processFile(c, "vaccine", comma)
	if err != nil {
		errors = append(errors, "Error processing 'vaccine': "+err.Error())
	}

	covidRaw, err := processFile(c, "covid", comma)
	if err != nil {
		errors = append(errors, "Error processing 'covid': "+err.Error())
	}

	covidGlobalRaw, err := processFile(c, "covid_global", commaCovidGlobal)
	if err != nil {
		warnings = append(warnings, "Warning: Unable to process the optional file 'covid_global'.: "+err.Error())
	}

	if len(errors) > 0 {
		c.JSON(422, gin.H{
			"message": "error on process one or more files",
			"errors": errors,
			"warning": warnings,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "success",
		"vaccinationRaw": vaccinationRaw,
		"vaccineRaw":     vaccineRaw,
		"covidRaw":       covidRaw,
		"covidGlobalRaw": covidGlobalRaw,
		"warning":        warnings,
	})
}

func processFile(c *gin.Context, key string, delimiter rune) (interface{}, error) {
	file, _, err := c.Request.FormFile(key)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rawData, err := service.ProcessCSV(file, delimiter)
	if err != nil {
		return nil, err
	}
	return rawData, nil
}