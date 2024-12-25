package handler

import (
	"encoding/csv"
	"fmt"

	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

type fileConfig struct {
	Key   string
	Comma rune
	Model interface{}
}

func ProcessCSV(c *gin.Context) {
	filesConfig := []fileConfig{
		{Key: "vaccination", Comma: ';', Model: &[]model.Vaccination{}},
		{Key: "vaccine", Comma: ';', Model: &[]model.Vaccine{}},
		{Key: "covid_global", Comma: ',', Model: &[]model.CovidGlobal{}},
		{Key: "covid", Comma: ';', Model: &[]model.Covid{}},
	}

	var errors []string
	var warnings []string
	parsedData := make(map[string]interface{})

	for _, fileCfg := range filesConfig {
		fileData, err := processFile(c, fileCfg)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error processing file '%s': %v", fileCfg.Key, err))
			continue
		}

		parsedData[fileCfg.Key] = fileData
	}

	if len(errors) > 0 {
		c.JSON(422, gin.H{
			"message":  "error processing some files",
			"errors":   errors,
			"warnings": warnings,
			"data":     parsedData,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "success",
		"errors":   errors,
		"warnings": warnings,
		// "data":     parsedData, // TODO: Make this a query parameter.
	})
}

func processFile(c *gin.Context, cfg fileConfig) (interface{}, error) {
	file, _, err := c.Request.FormFile(cfg.Key)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve file: %w", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = cfg.Comma

	if err := gocsv.UnmarshalCSV(csvReader, cfg.Model); err != nil {
		return nil, fmt.Errorf("error unmarshalling CSV data: %w", err)
	}

	return cfg.Model, nil
}