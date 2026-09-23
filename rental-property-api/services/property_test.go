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

// func Test
