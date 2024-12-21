package utils

import (
	"encoding/csv"
	"fmt"
	"log"
)



// ExtractDataFromCSV extracts data from a CSV file and returns it as a map
func ExtractDataFromCSV(file *csv.Reader) ([]map[string]string, error) {
	header, err := extractHeaderFromCSV(file)
	if err != nil {
		log.Println("Error extracting header from CSV file: " + err.Error())
		return nil, err
	}

	body, err := extractBodyFromCSV(file)
	if err != nil {
		log.Println("Error extracting body from CSV file: " + err.Error())
		return nil, err
	}

	result := ParseCSVToMap(header, body)
	return result, nil
}

// ParseCSVToMap parses a CSV file into a map
func ParseCSVToMap(header []string, body [][]string) []map[string]string {
	var records []map[string]string
	for _, row := range body {
		record := make(map[string]string)
		for i, key := range header {
			if i < len(row) {
				record[key] = row[i]
			}
		}
		records = append(records, record)
	}
	return records
}

// extractHeaderFromCSV extracts the header from a CSV file
func extractHeaderFromCSV(file *csv.Reader) ([]string, error) {
	headerRaw, err := file.Read()
	if err != nil {
		fmt.Println("Error reading header from CSV file: " + err.Error())
		return nil, err
	}
	return headerRaw, nil
}

// extractBodyFromCSV extracts the body (rows) from a CSV file
func extractBodyFromCSV(file *csv.Reader) ([][]string, error) {
	var body [][]string
	for {
		row, err := file.Read()
		if err != nil {
			if err.Error() != "EOF" {
				return nil, err
			}
			break
		}
		body = append(body, row)
	}
	return body, nil
}