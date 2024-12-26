package service

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"github.com/JrSchmidtt/covid-19-neo4J/src/model"
	"github.com/JrSchmidtt/covid-19-neo4J/src/storage"
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
	Load(data map[string]interface{}) error
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

func (s *etlService) Transform(data map[string]interface{}) error {
	fmt.Println("Transforming data...")

	return nil
}

// The Load method receives the transformed data and saves it to the database neo4j.
func (s *etlService) Load(data map[string]interface{}) error {
	for range data {
		fmt.Println("Loading data...")
	}
	vaccination := data["vaccination"].(*[]model.Vaccination)
	for _, v := range *vaccination {
		country := model.Country{
			Name: v.Country,
			Code: v.ISO3,
			Region: v.WHORegion,
		}
		_, err := storage.CreateCountry(country)
		if err != nil {
			fmt.Println(err)
		}
		_, err = storage.CreateVaccination(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	vaccine := data["vaccine"].(*[]model.Vaccine)
	for _, v := range *vaccine {
		_, err := storage.CreateVaccine(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	covidGlobal := data["covid_global"].(*[]model.CovidGlobal)
	for _, c := range *covidGlobal {
		_, err := storage.CreateCovidGlobal(c)
		if err != nil {
			fmt.Println(err)
		}
	}

	covid := data["covid"].(*[]model.Covid)
	for _, c := range *covid {
		_, err := storage.CreateCovid(c)
		if err != nil {
			fmt.Println(err)
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
