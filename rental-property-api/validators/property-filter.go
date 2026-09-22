package validators

import (
	"rental-property-api/models"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

// constant littarels
var validFeeds = map[int]bool{11: true, 12: true, 22: true, 24: true}
var validPropertyTypes = map[string]bool{"Hotel": true, "House": true, "Apartment": true, "Villa": true, "Resort": true, "Hostel": true}

func ValidateAndParseFilter(c *beego.Controller) (*models.PropertyFilter, []string) {
	// filters and errors lists
	filter := &models.PropertyFilter{}
	var errs []string

	// validate min price
	value := c.GetString("min_price")
	if value != "" {
		val, err := strconv.ParseFloat(value, 64)
		if err == nil {
			filter.MinPrice = &val
		} else {
			errs = append(errs, "invalid min_price: must be a number")
		}
	}

	// validate max price
	value = c.GetString("max_price")
	if value != "" {
		val, err := strconv.ParseFloat(value, 64)
		if err == nil {
			filter.MaxPrice = &val
		} else {
			errs = append(errs, "invalid max_price: must be a number")
		}
	}

	// validate min star rating
	value = c.GetString("min_star_rating")
	if value != "" {
		val, err := strconv.Atoi(value)
		if err == nil {
			filter.MinStarRating = &val
		} else {
			errs = append(errs, "invalid min_star_rating: must be an integer")
		}
	}

	// validate min review score
	value = c.GetString("min_review_score")
	if value != "" {
		val, err := strconv.ParseFloat(value, 64)
		if err == nil {
			filter.MinReviewScore = &val
		} else {
			errs = append(errs, "invalid min_review_score: must be a number")
		}
	}

	// validate min reviews
	value = c.GetString("min_reviews")
	if value != "" {
		val, err := strconv.Atoi(value)
		if err == nil {
			filter.MinReviews = &val
		} else {
			errs = append(errs, "invalid min_reviews: must be an integer")
		}
	}

	// validate published or not
	value = c.GetString("published")
	if value != "" {
		val, err := strconv.ParseBool(value)
		if err == nil {
			filter.Published = &val
		} else {
			errs = append(errs, "invalid published: must be true or false")
		}
	}

	// property type validation using constants
	value = c.GetString("property_type")
	if value != "" {
		if validPropertyTypes[value] {
			filter.PropertyType = &value
		} else {
			errs = append(errs, "invalid property_type: must be one of Hotel, House, Apartment, Villa, Resort, Hostel")
		}
	}

	// feed type validation using constants
	value = c.GetString("feed")
	if value != "" {
		val, err := strconv.Atoi(value)
		if err == nil && validFeeds[val] {
			filter.Feed = &val
		} else {
			errs = append(errs, "invalid feed: must be one of 11, 12, 22, 24")
		}
	}

	// validate bedroom numbers
	value = c.GetString("min_bedroom")
	if value != "" {
		val, err := strconv.Atoi(value)
		if err == nil {
			filter.MinBedroom = &val
		} else {
			errs = append(errs, "invalid min_bedroom: must be an integer")
		}
	}

	value = c.GetString("amenities")
	if value != "" {
		filter.Amenities = strings.Split(value, ",")
	}

	value = c.GetString("limit")
	if value != "" {
		val, err := strconv.Atoi(value)
		if err == nil && val >= 0 {
			filter.Limit = &val
		} else {
			errs = append(errs, "invalid limit: must be a non-negative integer")
		}
	}

	if len(errs) > 0 {
		return nil, errs
	}
	return filter, nil
}
