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
		logs.Critical("Error : sourcepath is not configured in app.conf")
		return err
	}
	if sourcePath == "" {
		logs.Critical("Error : sourcepath is not configured in app.conf")
		return err
	}

	// read raw data using the source file path
	rawData, err := os.ReadFile(sourcePath)
	if err != nil {
		logs.Error("Error : failed to read sourcefile %v", err)
		return err
	}

	// parse  raw json data on struct instance
	var parsedData []models.RentalPropertiesDTO

	parseErr := json.Unmarshal([]byte(rawData), &parsedData)
	if parseErr != nil {
		logs.Error("Error : failed to parse raw json data : %v", err)
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
			logs.Error("Error : failed to parse categories for id %s: %v", parsedData[i].ID, err)
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

// Get all filtered properties
func GetFilteredProperties(allProperties []models.RentalPropertiesDTO, filter *models.PropertyFilter) models.PropertyResult {
	// here properties are filtered according to filter
	matched := make([]*models.RentalPropertiesDTO, 0)
	for i := range allProperties {
		if matchesFilter(&allProperties[i], filter) {
			matched = append(matched, &allProperties[i])
		}
	}

	// transform matched items structure
	items := make([]models.PropertyResponseDTO, 0, len(matched))
	for _, p := range matched {
		items = append(items, toResponse(p))
	}

	totalCount := len(items)

	// if limit is provided and valid then response will be sent in that way
	if filter.Limit != nil && *filter.Limit < len(items) {
		limited := make([]models.PropertyResponseDTO, 0, *filter.Limit)
		for i := 0; i < *filter.Limit; i++ {
			limited = append(limited, items[i])
		}
		items = limited
		if *filter.Limit >= 0 && *filter.Limit < totalCount {
			totalCount = *filter.Limit
		}
	}

	// final result
	return models.PropertyResult{
		Count: totalCount,
		Items: items,
	}
}

// Get A  Property By Property ID
func GetAPropertyByID(properties []models.RentalPropertiesDTO, id string) (*models.PropertyResponseDTO, bool) {
	for i := range properties {
		if properties[i].ID == id {
			resp := toResponse(&properties[i])
			return &resp, true
		}
	}
	return nil, false
}
