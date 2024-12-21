package service

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
)

func ProcessCSV(DataFile multipart.File, comma rune) ([]map[string]string, error) {
	reader := csv.NewReader(DataFile)
	reader.Comma = comma
	vaccinationRawData, err := utils.ExtractDataFromCSV(reader) ; if err != nil {
		return nil, err
	}
	return vaccinationRawData, nil
}