package services

import "rental-property-api/models"

func matchesFilter(property *models.RentalPropertiesDTO, filter *models.PropertyFilter) bool {
	switch {
	case filter.MinPrice != nil && property.USDPrice < *filter.MinPrice:
		return false
	case filter.MaxPrice != nil && property.USDPrice > *filter.MaxPrice:
		return false
	case filter.MinStarRating != nil && property.StarRating < *filter.MinStarRating:
		return false
	case filter.MinReviewScore != nil && property.ReviewScoreGeneral < *filter.MinReviewScore:
		return false
	case filter.MinReviews != nil && property.NumberOfReviews < *filter.MinReviews:
		return false
	case filter.Published != nil && property.Published != *filter.Published:
		return false
	case filter.PropertyType != nil && property.PropertyType != *filter.PropertyType:
		return false
	case filter.Feed != nil && property.Feed != *filter.Feed:
		return false
	case filter.MinBedroom != nil && property.BedroomCount < *filter.MinBedroom:
		return false
	case len(filter.Amenities) > 0 && !hasAnyAmenity(property.AmenityCategories, filter.Amenities):
		return false
	}
	return true
}

func hasAnyAmenity(has, want []string) bool {
	set := make(map[string]bool, len(has))
	for _, hs := range has {
		set[hs] = true
	}
	for _, wt := range want {
		if set[wt] {
			return true
		}
	}
	return false
}
