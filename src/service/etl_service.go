package service

import (
	"encoding/csv"
	"mime/multipart"
	"reflect"
	"strconv"
	"time"

	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/JrSchmidtt/covid-19-neo4J/src/utils"
)

func ExtractData(DataFile multipart.File, comma rune) ([]map[string]string, error) {
	reader := csv.NewReader(DataFile)
	reader.Comma = comma
	vaccinationRawData, err := utils.ExtractDataFromCSV(reader)
	if err != nil {
		return nil, err
	}
	return vaccinationRawData, nil
}

func TransformVaccination(rawData []map[string]string) ([]model.Vaccination, error) {
	var vaccinationData []model.Vaccination

	for _, data := range rawData {
		var vaccination model.Vaccination
		v := reflect.ValueOf(&vaccination).Elem()
		t := v.Type()

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldName := field.Tag.Get("csv")
			if fieldName == "" {
				fieldName = field.Name // Use the field name if the csv tag is not defined
			}

			value, ok := data[fieldName]
			if !ok {
				continue // Skip if the field is not present in the data
			}

			fieldValue := v.Field(i)
			if !fieldValue.CanSet() {
				continue
			}

			switch fieldValue.Kind() {
			case reflect.Float64:
				parsedValue, err := strconv.ParseFloat(value, 64)
				if err != nil {
					parsedValue = 0.0 // Set to zero value for float64
				}
				fieldValue.SetFloat(parsedValue)

			case reflect.Int64:
				// Try to parse to int64, treating values ​​in scientific notation
				parsedFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					parsedFloat = 0.0 // Set to zero value for float64
				}
				fieldValue.SetInt(int64(parsedFloat)) // Convert to int64
			case reflect.Struct: // Type is time.Time
				if fieldType := field.Type; fieldType == reflect.TypeOf(time.Time{}) {
					parsedDate := utils.ParseDate(value)
					fieldValue.Set(reflect.ValueOf(parsedDate))
				}

			case reflect.String:
				fieldValue.SetString(value)

			default:
				continue
			}
		}

		vaccinationData = append(vaccinationData, vaccination)
	}

	return vaccinationData, nil
}