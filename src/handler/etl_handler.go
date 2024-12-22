package handler

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"github.com/JrSchmidtt/covid-19-neo4J/src/service"
	"github.com/gin-gonic/gin"
)

const comma = ';'
const commaCovidGlobal = ','
type filesConfig struct {
	key       string
	delimiter rune
	isOptional bool
}

func ProcessCSV(c *gin.Context) {
	var errors []string
	var warnings []string

	files := make(map[string][]byte)
	filesConfig := []filesConfig{
		{"vaccination", comma, false},
		{"vaccine", comma, false},
		{"covid", comma, false},
		{"covid_global", commaCovidGlobal, true},
	}

	for _, config := range filesConfig {
		file, _, err := c.Request.FormFile(config.key)
		if err != nil {
			if config.isOptional {
				warnings = append(warnings, "Warning: Unable to process the optional file '"+config.key+"': "+err.Error())
			} else {
				errors = append(errors, "Error retrieving file '"+config.key+"': "+err.Error())
			}
			continue
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			if config.isOptional {
				warnings = append(warnings, "Warning: Unable to read the optional file '"+config.key+"': "+err.Error())
			} else {
				errors = append(errors, "Error reading file '"+config.key+"': "+err.Error())
			}
			continue
		}
		files[config.key] = data
	}

	if len(errors) > 0 {
		c.JSON(422, gin.H{
			"message": "error on process one or more files",
			"errors":  errors,
			"warning": warnings,
		})
		return
	}

	// Process files concurrently
	results := make(map[string]interface{})
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(files))
	for key, data := range files {
		go func(fileKey string, fileData []byte, delimiter rune) {
			defer wg.Done()
			rawData, err := service.ExtractData(bytes.NewReader(fileData), delimiter)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				errors = append(errors, "Error processing '"+fileKey+"': "+err.Error())
				return
			}
			results[fileKey] = rawData
		}(key, data, getDelimiterForKey(key))
	}
	wg.Wait()

	vacination, err := service.TransformVaccination(results["vaccination"].([]map[string]string))
	if err != nil {
		errors = append(errors, "Error transforming vaccination data: "+err.Error())
	}


	if len(errors) > 0 {
		c.JSON(422, gin.H{
			"message": "error on process one or more files",
			"errors":  errors,
			"warnings": warnings,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		// "results": results,
		"warning": warnings,
		"vaccination": vacination,
	})
}

func getDelimiterForKey(key string) rune {
	if key == "covid_global" {
		return commaCovidGlobal
	}
	return comma
}