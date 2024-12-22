package handler

import (
	"encoding/csv"
	"fmt"

	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

func ProcessCSV(c *gin.Context) {
    var errors []string
    var warnings []string

    // Abre o arquivo CSV enviado
    file, _, err := c.Request.FormFile("vaccination")
    if err != nil {
        errors = append(errors, "Error retrieving file 'vaccination': "+err.Error())
        c.JSON(422, gin.H{
            "message":  "error processing file",
            "errors":   errors,
            "warnings": warnings,
        })
        return
    }
    defer file.Close()

    // Configura o delimitador como ponto e vírgula
    csvReader := csv.NewReader(file)
    csvReader.Comma = ';'

    var vaccinationData []model.Vaccination
    if err := gocsv.UnmarshalCSV(csvReader, &vaccinationData); err != nil {
        errors = append(errors, "Error unmarshalling CSV data: "+err.Error())
		for _, v := range vaccinationData {
			fmt.Println(v)
		}
        c.JSON(422, gin.H{
            "message":  "error parsing CSV data",
            "errors":   errors,
            "warnings": warnings,
        })
        return
    }

    // Exibe os dados processados para depuração
    for _, v := range vaccinationData {
        fmt.Println(v)
    }

    c.JSON(200, gin.H{
        "message":  "success",
        "errors":   errors,
        "warnings": warnings,
        "data":     vaccinationData,
    })
}