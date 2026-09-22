package models

type PropertyResponseDTO struct {
	ID        string       `json:"ID"`
	Feed      int          `json:"Feed"`
	Published bool         `json:"Published"`
	GeoInfo   GeoInfo      `json:"GeoInfo"`
	Property  PropertyInfo `json:"Property"`
}

type GeoInfo struct {
	Breadcrumbs []string `json:"Breadcrumbs"`
	City        string   `json:"City"`
	Country     string   `json:"Country"`
	CountryCode string   `json:"CountryCode"`
	Name        string   `json:"Name"`
	LocationID  string   `json:"LocationID"`
	Lat         float64  `json:"Lat"`
	Lon         float64  `json:"Lon"`
	State       string   `json:"State"`
	StateAbbr   string   `json:"StateAbbr"`
}

type PropertyInfo struct {
	Amenities    []string  `json:"Amenities"`
	Name         string    `json:"Name"`
	Slug         string    `json:"Slug"`
	PropertyType string    `json:"PropertyType"`
	Price        float64   `json:"Price"`
	ReviewScore  float64   `json:"ReviewScore"`
	StarRating   int       `json:"StarRating"`
	Counts       Counts    `json:"Counts"`
	Image        ImageInfo `json:"Image"`
}

type Counts struct {
	Bathroom  int `json:"Bathroom"`
	Bedroom   int `json:"Bedroom"`
	Reviews   int `json:"Reviews"`
	Occupancy int `json:"Occupancy"`
}

type ImageInfo struct {
	Count  int      `json:"Count"`
	Images []string `json:"Images"`
}

// List-wrapper types for GET /v1/properties
type PropertyListResponse struct {
	Result PropertyResult `json:"Result"`
}

type PropertyResult struct {
	Count int                   `json:"Count"`
	Items []PropertyResponseDTO `json:"Items"`
}
