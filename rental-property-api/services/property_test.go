package services

import (
	"reflect"
	"rental-property-api/models"
	"testing"
)

// requirement : 1 (Transform)

func TestToResponse(t *testing.T) {
	cases := []struct {
		name            string
		input           models.RentalPropertiesDTO
		wantLAT         float64
		wantLON         float64
		wantImgCnt      int
		wantBreadCrumbs []string
	}{
		{
			name: "property with multiple images and breadcrumbs",
			input: models.RentalPropertiesDTO{
				ID:           "p1",
				PropertyName: "Sunny Villa",
				Images:       []string{"img1.jpg", "img2.jpg", "img3.jpg"},
				LonLat: models.LonLat{
					Coordinates: []float64{-122.4194, 37.7749},
				},
				Breadcrumbs: []string{
					"United States",
					"California",
					"San Francisco",
				},
			},
			wantLAT:         37.7749,
			wantLON:         -122.4194,
			wantImgCnt:      3,
			wantBreadCrumbs: []string{"United States", "California", "San Francisco"},
		},

		{
			name: "property with one image and two breadcrumbs",
			input: models.RentalPropertiesDTO{
				ID:           "p2",
				PropertyName: "Tokyo Apartment",
				Images:       []string{"tokyo.jpg"},
				LonLat: models.LonLat{
					Coordinates: []float64{139.6917, 35.6895},
				},
				Breadcrumbs: []string{
					"Japan",
					"Tokyo",
				},
			},
			wantLAT:         35.6895,
			wantLON:         139.6917,
			wantImgCnt:      1,
			wantBreadCrumbs: []string{"Japan", "Tokyo"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := toResponse(&tc.input)

			if res.Property.Image.Count != tc.wantImgCnt {
				t.Errorf("Failed : case-image-coount = %d and want-image-count = %d", res.Property.Image.Count, tc.wantImgCnt)
			}
			if res.GeoInfo.Lat != tc.wantLAT {
				t.Errorf("Failed : case-geoinfo-lat = %f and want-lat = %f", res.GeoInfo.Lat, tc.wantLAT)
			}
			if res.GeoInfo.Lon != tc.wantLON {
				t.Errorf("Failed : case-geoinfo-lon = %f and want-Lon = %f", res.GeoInfo.Lon, tc.wantLON)
			}
			if !reflect.DeepEqual(res.GeoInfo.Breadcrumbs, tc.wantBreadCrumbs) {
				t.Errorf("Failed : case-breadcrumbs = %v, and want-breadcrumbs = %v", res.GeoInfo.Breadcrumbs, tc.wantBreadCrumbs)
			}
		})
	}
}

// requirement : 2 (Filter AND)

// helper function for filter options
func floatPtr(v float64) *float64 { return &v }
func intPtr(v int) *int           { return &v }
func strPtr(v string) *string     { return &v }
func boolPtr(v bool) *bool        { return &v }

// test properties for testing filter options
var testProperties = []models.RentalPropertiesDTO{
	{
		ID:                 "p1",
		Feed:               11,
		Country:            "United States",
		CountryCode:        "US",
		State:              "California",
		StateAbbr:          "CA",
		City:               "San Francisco",
		Display:            "Sunny Villa, San Francisco",
		LocationID:         "loc-001",
		PropertyName:       "Sunny Villa",
		PropertySlug:       "sunny-villa-san-francisco",
		PropertyType:       "villa",
		USDPrice:           250.00,
		Occupancy:          4,
		BedroomCount:       2,
		BathroomCount:      2,
		NumberOfReviews:    128,
		ReviewScoreGeneral: 4.8,
		StarRating:         5,
		AmenityCategories:  []string{"pool", "wifi", "parking"},
		LonLat: models.LonLat{
			Coordinates: []float64{-122.4194, 37.7749}, // [lon, lat]
		},
		Categories:  "Vacation Rentals > Villas",
		Breadcrumbs: []string{"United States", "California", "San Francisco"},
		Published:   true,
		Images:      []string{"villa-1.jpg", "villa-2.jpg", "villa-3.jpg"},
	},

	{
		ID:                 "p2",
		Feed:               11,
		Country:            "United States",
		CountryCode:        "US",
		State:              "New York",
		StateAbbr:          "NY",
		City:               "New York",
		Display:            "Modern Apartment, New York",
		LocationID:         "loc-002",
		PropertyName:       "Modern Manhattan Apartment",
		PropertySlug:       "modern-manhattan-apartment",
		PropertyType:       "apartment",
		USDPrice:           180.00,
		Occupancy:          2,
		BedroomCount:       1,
		BathroomCount:      1,
		NumberOfReviews:    75,
		ReviewScoreGeneral: 4.5,
		StarRating:         4,
		AmenityCategories:  []string{"wifi", "gym"},
		LonLat: models.LonLat{
			Coordinates: []float64{-74.0060, 40.7128}, // [lon, lat]
		},
		Categories:  "Vacation Rentals > Apartments",
		Breadcrumbs: []string{"United States", "New York", "New York"},
		Published:   false,
		Images:      []string{"apartment-1.jpg", "apartment-2.jpg"},
	},

	{
		ID:                 "p3",
		Feed:               12,
		Country:            "Japan",
		CountryCode:        "JP",
		State:              "Tokyo",
		StateAbbr:          "TK",
		City:               "Tokyo",
		Display:            "Traditional House, Tokyo",
		LocationID:         "loc-003",
		PropertyName:       "Tokyo Traditional House",
		PropertySlug:       "tokyo-traditional-house",
		PropertyType:       "house",
		USDPrice:           320.00,
		Occupancy:          6,
		BedroomCount:       3,
		BathroomCount:      2,
		NumberOfReviews:    210,
		ReviewScoreGeneral: 4.9,
		StarRating:         5,
		AmenityCategories:  []string{"wifi", "parking", "garden"},
		LonLat: models.LonLat{
			Coordinates: []float64{139.6917, 35.6895}, // [lon, lat]
		},
		Categories:  "Vacation Rentals > Houses",
		Breadcrumbs: []string{"Japan", "Tokyo"},
		Published:   true,
		Images:      []string{"house-1.jpg", "house-2.jpg", "house-3.jpg", "house-4.jpg"},
	},

	{
		ID:                 "p4",
		Feed:               22,
		Country:            "Bangladesh",
		CountryCode:        "BD",
		State:              "Dhaka",
		StateAbbr:          "DH",
		City:               "Dhaka",
		Display:            "Cozy Apartment, Dhaka",
		LocationID:         "loc-004",
		PropertyName:       "Cozy Dhaka Apartment",
		PropertySlug:       "cozy-dhaka-apartment",
		PropertyType:       "apartment",
		USDPrice:           75.00,
		Occupancy:          3,
		BedroomCount:       2,
		BathroomCount:      1,
		NumberOfReviews:    42,
		ReviewScoreGeneral: 4.2,
		StarRating:         3,
		AmenityCategories:  []string{"pool", "wifi"},
		LonLat: models.LonLat{
			Coordinates: []float64{90.4125, 23.8103}, // [lon, lat]
		},
		Categories:  "Vacation Rentals > Apartments",
		Breadcrumbs: []string{"Bangladesh", "Dhaka"},
		Published:   true,
		Images:      []string{"dhaka-1.jpg"},
	},
}

// helper get ids from the filtered properties
func getPropertyIDs(properties []models.PropertyResponseDTO) []string {
	var IDs []string
	for i := 0; i < len(properties); i++ {
		IDs = append(IDs, properties[i].ID)
	}
	return IDs
}

func TestGetFilteredPropertiesAnd(t *testing.T) {
	cases := []struct {
		name      string
		input     *models.PropertyFilter
		wantIDs   []string
		wantCount int
	}{
		{
			name: "minimum price",
			input: &models.PropertyFilter{
				MinPrice: floatPtr(150),
			},
			wantIDs: []string{"p1", "p2", "p3"},
		},
		{
			name: "maximum price",
			input: &models.PropertyFilter{
				MaxPrice: floatPtr(100),
			},
			wantIDs: []string{"p4"},
		},
		{
			name: "price range",
			input: &models.PropertyFilter{
				MinPrice: floatPtr(100),
				MaxPrice: floatPtr(200),
			},
			wantIDs: []string{"p2"},
		},
		{
			name: "minimum star rating",
			input: &models.PropertyFilter{
				MinStarRating: intPtr(4),
			},
			wantIDs: []string{"p1", "p2", "p3"},
		},
		{
			name: "minimum review score",
			input: &models.PropertyFilter{
				MinReviewScore: floatPtr(4.5),
			},
			wantIDs: []string{"p1", "p2", "p3"},
		},
		{
			name: "minimum reviews",
			input: &models.PropertyFilter{
				MinReviews: intPtr(100),
			},
			wantIDs: []string{"p1", "p3"},
		},
		{
			name: "published properties",
			input: &models.PropertyFilter{
				Published: boolPtr(true),
			},
			wantIDs: []string{"p1", "p3", "p4"},
		},
		{
			name: "property type",
			input: &models.PropertyFilter{
				PropertyType: strPtr("villa"),
			},
			wantIDs: []string{"p1"},
		},
		{
			name: "feed",
			input: &models.PropertyFilter{
				Feed: intPtr(11),
			},
			wantIDs: []string{"p1", "p2"},
		},
		{
			name: "minimum bedrooms",
			input: &models.PropertyFilter{
				MinBedroom: intPtr(2),
			},
			wantIDs: []string{"p1", "p3", "p4"},
		},
		{
			name: "amenities OR",
			input: &models.PropertyFilter{
				Amenities: []string{"pool", "parking"},
			},
			wantIDs: []string{"p1", "p3", "p4"},
		},
		{
			name: "limit",
			input: &models.PropertyFilter{
				Limit: intPtr(2),
			},
			wantIDs:   []string{"p1", "p2"},
			wantCount: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := GetFilteredProperties(testProperties, tc.input)
			if !reflect.DeepEqual(getPropertyIDs(res.Items), tc.wantIDs) {
				t.Errorf("filtered-property-ids = %v, wanted-property-ids %v", getPropertyIDs(res.Items), tc.wantIDs)
			}
			wantCount := tc.wantCount
			if wantCount == 0 {
				wantCount = len(tc.wantIDs)
			}
			if res.Count != wantCount {
				t.Errorf("filtered-property-count = %d, wanted-property-count %d", res.Count, wantCount)
			}
		})
	}
}

// requirement : 3 (Filter amenities OR)

func TestGetFilteredPropertiesAmenitiesOR(t *testing.T) {
	cases := []struct {
		name    string
		input   *models.PropertyFilter
		wantIDs []string
	}{
		{
			name: "amenities OR - single amenity match",
			input: &models.PropertyFilter{
				Amenities: []string{"gym"},
			},
			wantIDs: []string{"p2"},
		},
		{
			name: "amenities OR - multiple amenities match different properties",
			input: &models.PropertyFilter{
				Amenities: []string{"pool", "garden"},
			},
			wantIDs: []string{"p1", "p3", "p4"},
		},
		{
			name: "amenities OR - no property has any requested amenity",
			input: &models.PropertyFilter{
				Amenities: []string{"sauna", "hot tub"},
			},
			wantIDs: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := GetFilteredProperties(testProperties, tc.input)
			if !reflect.DeepEqual(getPropertyIDs(res.Items), tc.wantIDs) {
				t.Errorf("filtered-property-ids = %v, wanted-property-ids %v", getPropertyIDs(res.Items), tc.wantIDs)
			}
		})
	}
}

// requirement : 4 (Filter - Combined AND, Amenities)
func TestGetFilteredPropertiesCombined(t *testing.T) {

	cases := []struct {
		name    string
		input   *models.PropertyFilter
		wantIDs []string
	}{
		{
			name: "feed AND published AND amenities OR",
			input: &models.PropertyFilter{
				Feed:      intPtr(11),
				Published: boolPtr(true),
				Amenities: []string{"pool", "wifi"},
			},
			// p1: feed 11, published true, has pool and wifi matches
			wantIDs: []string{"p1"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := GetFilteredProperties(testProperties, tc.input)
			if !reflect.DeepEqual(getPropertyIDs(res.Items), tc.wantIDs) {
				t.Errorf("filtered-property-ids = %v, wanted-property-ids %v", getPropertyIDs(res.Items), tc.wantIDs)
			}
		})
	}
}

// requirement : 5 (Filter - empty values)
func TestGetFilteredProperties_EmptyResult(t *testing.T) {
	cases := []struct {
		name   string
		filter *models.PropertyFilter
	}{
		{
			name: "no property satisfies min_price",
			filter: &models.PropertyFilter{
				MinPrice: floatPtr(9999),
			},
			// highest price in testProperties is p3 at $320 -> no match
		},
		{
			name: "AND combination no property satisfies",
			filter: &models.PropertyFilter{
				Feed:         intPtr(11),
				PropertyType: strPtr("house"),
			},
			// feed 11 properties are p1 (villa) and p2 (apartment); neither is "house" -> no match
		},
		{
			name: "amenities OR - none have any requested amenity",
			filter: &models.PropertyFilter{
				Amenities: []string{"sauna", "hot tub"},
			},
			// none of p1-p4 list sauna or hot tub -> no match
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetFilteredProperties(testProperties, tc.filter)

			if got.Items == nil {
				t.Errorf("Items = nil, want non-nil empty slice")
			}
			if len(got.Items) != 0 {
				t.Errorf("len(Items) = %d, want 0", len(got.Items))
			}
			if got.Count != 0 {
				t.Errorf("Count = %d, want 0", got.Count)
			}
		})
	}
}

// requirement : 6 (Filter - get by id)
func TestGetAPropertyByID(t *testing.T) {
	cases := []struct {
		name      string
		id        string
		wantFound bool
	}{
		{
			name:      "found",
			id:        "p3",
			wantFound: true,
		},
		{
			name:      "not found",
			id:        "does-not-exist",
			wantFound: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, found := GetAPropertyByID(testProperties, tc.id)
			if found != tc.wantFound {
				t.Fatalf("found = %v, want %v", found, tc.wantFound)
			}
			if tc.wantFound {
				if res == nil {
					t.Fatal("expected non-nil result when found")
				}
				if res.ID != tc.id {
					t.Errorf("ID = %s, want %s", res.ID, tc.id)
				}
			} else if res != nil {
				t.Errorf("expected nil result when not found, res %+v", res)
			}
		})
	}
}
