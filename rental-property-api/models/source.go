package models

type RentalPropertiesDTO struct {
	ID                 string   `json:"id"`
	Feed               int      `json:"feed"`
	Country            string   `json:"country"`
	CountryCode        string   `json:"country_code"`
	State              string   `json:"state"`
	StateAbbr          string   `json:"state_abbr"`
	City               string   `json:"city"`
	Display            string   `json:"display"`
	LocationID         string   `json:"location_id"`
	PropertyName       string   `json:"property_name"`
	PropertySlug       string   `json:"property_slug"`
	PropertyType       string   `json:"property_type_category"`
	USDPrice           float64  `json:"usd_price"`
	Occupancy          int      `json:"occupancy"`
	BedroomCount       int      `json:"bedroom_count"`
	BathroomCount      int      `json:"bathroom_count"`
	NumberOfReviews    int      `json:"number_of_review"`
	ReviewScoreGeneral float64  `json:"review_score_general"`
	StarRating         int      `json:"star_rating"`
	AmenityCategories  []string `json:"amenity_categories"`
	LonLat             LonLat   `json:"lonlat"`     // created a new type for nested key-value
	Categories         string   `json:"categories"` //this field must be parsed as soon as the data is loaded in the memory
	Breadcrumbs        []string `json:"-"`
	Published          bool     `json:"published"`
	Images             []string `json:"images"`
}

// Nested key-value type
type LonLat struct {
	Coordinates []float64 `json:"coordinates"`
}

// for parcing categoris field and store in breadcrumbs
type CategoryEntry struct {
	LocationID string   `json:"LocationID"`
	Name       string   `json:"Name"`
	Type       string   `json:"Type"`
	Slug       string   `json:"Slug"`
	Display    []string `json:"Display"`
}
