package utils

import (
	"encoding/csv"
	"log"
	"strings"
)

// ExtractDataFromCSV extracts data from a CSV file and returns it as a map
func ExtractDataFromCSV(file *csv.Reader, separator string) ([]map[string]string, error) {
	header, err := extractHeaderFromCSV(file, separator)
	if err != nil {
		log.Println("Error extracting header from CSV file" + err.Error())
		return nil, err
	}
	body, err := extractBodyFromCSV(file)
	if err != nil {
		log.Println("Error extracting body from CSV file" + err.Error())
		return nil, err
	}
	result := ParseCSVToMap(header, body)
	return result, nil
}

// ParseCSVToMap parses a CSV file into a map
func ParseCSVToMap(header []string, body []string) []map[string]string {
	var records []map[string]string
	for i := 0; i < len(body); i += len(header) {
		record := make(map[string]string)
		for j, key := range header {
			if i+j < len(body) {
				record[key] = body[i+j]
			}
		}
		records = append(records, record)
	}
	return records
}

// extractHeaderFromCSV extracts the header from a CSV file
func extractHeaderFromCSV(file *csv.Reader, separator string) ([]string, error) {
	header := []string{}
	headerRaw, err := file.Read()
	if err != nil {
		return nil, err
	}
	for _, value := range headerRaw {
		splitValues := strings.Split(value, separator)
		header = append(header, splitValues...)
	}
	return header, nil
}

func extractBodyFromCSV(file *csv.Reader) ([]string, error) {
	body := []string{}
	bodyRaw, err := file.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, row := range bodyRaw {
		for _, value := range row {
			splitValues := strings.Split(value, ";")
			body = append(body, splitValues...)
		}
	}
	return body, nil
}