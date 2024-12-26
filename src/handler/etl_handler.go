package handler

import (
	"mime/multipart"
	"net/http"

	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/JrSchmidtt/covid-19-neo4J/src/service"
	"github.com/gin-gonic/gin"
)


func ProcessCSV(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse form data",
			"error":   err.Error(),
		})
		return
	}

	files := make(map[string]*multipart.FileHeader)
	for key, fileHeaders := range form.File {
		if len(fileHeaders) > 0 {
			files[key] = fileHeaders[0]
		}
	}

	etlService := service.NewETLService()
	etlService.SetFileConfig([]service.FileConfig{
		{Key: "vaccination", Comma: ';', Model: &[]model.Vaccination{}},
		{Key: "vaccine", Comma: ';', Model: &[]model.Vaccine{}},
		{Key: "covid_global", Comma: ',', Model: &[]model.CovidGlobal{}},
		{Key: "covid", Comma: ';', Model: &[]model.Covid{}},
	})
	
	parsedData, errors, warnings, err := etlService.Extract(files)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message":  "error processing files",
			"errors":   errors,
			"warnings": warnings,
			"data":     parsedData,
		})
		return
	}
	if err := etlService.Transform(parsedData); err != nil {
		c.JSON(500, gin.H{
			"message": "error transforming data",
			"error":   err.Error(),
		})
		return
	}
	// vaccination := parsedData["vaccination"].(*[]model.Vaccination)
	// vaccine := parsedData["vaccine"].(*[]model.Vaccine)
	// covidGlobal := parsedData["covid_global"].(*[]model.CovidGlobal)
	// covid := parsedData["covid"].(*[]model.Covid)

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"errors":   errors,
		"warnings": warnings,
		 "data":     parsedData,
		// "data":     parsedData, // TODO: Make this a query parameter.
	})
}