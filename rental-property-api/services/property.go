package services

import (
	"encoding/json"
	"os"
	"rental-property-api/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type propertyStore struct {
	isLoaded bool
	data     []models.RentalPropertiesDTO
}

var propertyInstance propertyStore

func LoadData() error {

	// extract the source file path
	sourcePath, err := web.AppConfig.String("sourcepath")

	// check if path is valid or not
	if err != nil {
		logs.Critical("sourcepath is not configured in app.conf")
		return err
	}
	if sourcePath == "" {
		logs.Critical("sourcepath is not configured in app.conf")
		return err
	}

	// read raw data using the source file path
	rawData, err := os.ReadFile(sourcePath)
	if err != nil {
		logs.Error("failed to read sourcefile %v", err)
		return err
	}

	// parse  raw json data on struct instance
	var parsedData []models.RentalPropertiesDTO

	parseErr := json.Unmarshal([]byte(rawData), &parsedData)
	if parseErr != nil {
		logs.Error("failed to parse raw json data : %v", err)
		return err
	}

	// before loading the parsed data , categories field need to be parsed at load time to avoid redundancy
	for i := range parsedData {
		if parsedData[i].Categories == "" {
			continue
		}

		var categoryEntries []models.CategoryEntry
		err := json.Unmarshal([]byte(parsedData[i].Categories), &categoryEntries)
		if err != nil {
			logs.Error("failed to parse categories for id %s: %v", parsedData[i].ID, err)
			continue
		}

		breadcrumbs := make([]string, 0, len(categoryEntries))
		for _, entry := range categoryEntries {
			breadcrumbs = append(breadcrumbs, entry.Name)
		}
		parsedData[i].Breadcrumbs = breadcrumbs
	}

	// load data on memory

	propertyInstance.data = parsedData
	propertyInstance.isLoaded = true

	return nil
}

// singleton object model and only way to get propertyStore instance
// ensures both once data loading and encapsulation properties
func GetPropertyStore() (*propertyStore, error) {
	if !propertyInstance.isLoaded {
		err := LoadData()
		if err != nil {
			return nil, err
		}
	}
	return &propertyInstance, nil
}

// Get data from property store
func (p *propertyStore) GetData() []models.RentalPropertiesDTO {
	return p.data
}
