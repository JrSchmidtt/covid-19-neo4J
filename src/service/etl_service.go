package service

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"reflect"
	"regexp"
	"time"

	"github.com/gocarina/gocsv"
)

type FileConfig struct {
	Key   string
	Comma rune
	Model interface{}
}

// Extract, Transform and Load (ETL) service interface.
// It defines the methods that must be implemented by the ETL service.
type ETLService interface {
	SetFileConfig(configs []FileConfig)
	Extract(files map[string]*multipart.FileHeader) (map[string]interface{}, []string, []string, error)
	Transform(data map[string]interface{}) error
}

type etlService struct {
	filesConfig []FileConfig
}

func NewETLService() ETLService {
	return &etlService{}
}

// SetFileConfig sets the configuration for the files that will be processed by the service.
func (s *etlService) SetFileConfig(configs []FileConfig) {
	s.filesConfig = configs
}

// The Extract method reads the files from the request and processes them according to the configuration
// provided in the service. It returns the parsed data, a list of errors and a list of warnings.
func (s *etlService) Extract(files map[string]*multipart.FileHeader) (map[string]interface{}, []string, []string, error) {
	var errors []string
	var warnings []string
	parsedData := make(map[string]interface{})

	for _, fileCfg := range s.filesConfig {
		fileHeader, ok := files[fileCfg.Key]
		if !ok {
			warnings = append(warnings, fmt.Sprintf("File '%s' not provided", fileCfg.Key))
			continue
		}

		fileData, err := processFile(fileHeader, fileCfg)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Error processing file '%s': %v", fileCfg.Key, err))
			continue
		}

		parsedData[fileCfg.Key] = fileData
	}

	if len(errors) > 0 {
		return parsedData, errors, warnings, fmt.Errorf("errors occurred while processing files")
	}

	return parsedData, errors, warnings, nil
}

// The Transform method receives the parsed data and applies transformations to it.
// In this case, it parses the date fields in the data to the time.Time type.
// The date fields are identified by a regular expression that matches the pattern dd/MM/yyyy.
func (s *etlService) Transform(data map[string]interface{}) error {
    // Regex patterns to match dates in DD/MM/YYYY and YYYY-MM-DD formats
    dateRegexDDMMYYYY := regexp.MustCompile(`\d{2}/\d{2}/\d{4}`)
    dateRegexYYYYMMDD := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

    for key, rawData := range data {
        // Check if the value is a pointer to a slice
        value := reflect.ValueOf(rawData)
        if value.Kind() == reflect.Ptr {
            value = value.Elem() // Dereference the pointer
        }

        if value.Kind() != reflect.Slice {
            return fmt.Errorf("expected slice for key '%s', got %T", key, rawData)
        }

        // Iterate over each item in the slice
        for i := 0; i < value.Len(); i++ {
            item := value.Index(i)
            if item.Kind() == reflect.Ptr {
                item = item.Elem() // Dereference the pointer (struct)
            }

            if item.Kind() != reflect.Struct {
                return fmt.Errorf("expected struct in slice for key '%s', got %T", key, item.Interface())
            }

            // Iterate over each field in the struct
            for j := 0; j < item.NumField(); j++ {
                field := item.Field(j)
                fieldType := item.Type().Field(j)

                // Check if the field is exported and a string
                if field.Kind() == reflect.String && field.CanSet() {
                    strValue := field.String()

                    var parsedDate time.Time
                    var err error

                    // Check and process date in DD/MM/YYYY format
                    if dateRegexDDMMYYYY.MatchString(strValue) {
                        parsedDate, err = time.Parse("02/01/2006", strValue)
                        if err != nil {
                            return fmt.Errorf("failed to parse date in field '%s' for key '%s': %v", fieldType.Name, key, err)
                        }
                    }

                    // Check and process date in YYYY-MM-DD format
                    if dateRegexYYYYMMDD.MatchString(strValue) {
                        parsedDate, err = time.Parse("2006-01-02", strValue)
                        if err != nil {
                            return fmt.Errorf("failed to parse date in field '%s' for key '%s': %v", fieldType.Name, key, err)
                        }
                    }

                    // If a date was successfully parsed, save it as an ISO 8601 string with a timestamp
                    if !parsedDate.IsZero() {
                        field.SetString(parsedDate.Format(time.RFC3339)) // Save as ISO 8601 format
                    }
                }
            }
        }
    }
    return nil
}

func processFile(fileHeader *multipart.FileHeader, cfg FileConfig) (interface{}, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = cfg.Comma

	if err := gocsv.UnmarshalCSV(csvReader, cfg.Model); err != nil {
		return nil, fmt.Errorf("error unmarshalling CSV data: %w", err)
	}

	return cfg.Model, nil
}