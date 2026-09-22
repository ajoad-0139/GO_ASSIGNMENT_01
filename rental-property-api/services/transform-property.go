package services

import (
	"rental-property-api/models"
)

func toResponse(p *models.RentalPropertiesDTO) models.PropertyResponseDTO {
	var lat, lon float64
	if len(p.LonLat.Coordinates) >= 2 {
		lon = p.LonLat.Coordinates[0]
		lat = p.LonLat.Coordinates[1]
	}

	return models.PropertyResponseDTO{
		ID:        p.ID,
		Feed:      p.Feed,
		Published: p.Published,
		GeoInfo: models.GeoInfo{
			City:        p.City,
			Breadcrumbs: p.Breadcrumbs,
			Country:     p.Country,
			Name:        p.Display,
			CountryCode: p.CountryCode,
			LocationID:  p.LocationID,
			Lat:         lat,
			Lon:         lon,
			StateAbbr:   p.StateAbbr,
			State:       p.State,
		},
		Property: models.PropertyInfo{
			Slug:         p.PropertySlug,
			Name:         p.PropertyName,
			Price:        p.USDPrice,
			StarRating:   p.StarRating,
			PropertyType: p.PropertyType,
			Amenities:    p.AmenityCategories,
			ReviewScore:  p.ReviewScoreGeneral,
			Counts: models.Counts{
				Bathroom:  p.BathroomCount,
				Bedroom:   p.BedroomCount,
				Reviews:   p.NumberOfReviews,
				Occupancy: p.Occupancy,
			},
			Image: models.ImageInfo{
				Count:  len(p.Images),
				Images: p.Images,
			},
		},
	}
}
