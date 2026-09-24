package controllers

import (
	"rental-property-api/models"
	"rental-property-api/services"
	"rental-property-api/utils"
	"rental-property-api/validators"

	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

// @Title Get All Properties
// @Description Returns a filtered list of rental properties. All query params are optional and can be combined.
// @Param   min_price         query   string  false   "Minimum price (USD)"
// @Param   max_price         query   string  false   "Maximum price (USD)"
// @Param   min_star_rating   query   int     false   "Minimum star rating"
// @Param   min_review_score  query   string  false   "Minimum review score"
// @Param   min_reviews       query   int     false   "Minimum number of reviews"
// @Param   published         query   bool    false   "Filter by published status"
// @Param   property_type     query   string  false   "One of Hotel, House, Apartment, Villa, Resort, Hostel"
// @Param   feed              query   int     false   "One of 11, 12, 22, 24"
// @Param   min_bedroom       query   int     false   "Minimum number of bedrooms"
// @Param   amenities         query   string  false   "Comma-separated list, e.g. wifi,pool"
// @Param   limit             query   int     false   "Maximum number of results to return"
// @Success 200 {object} models.PropertyListResponse
// @Failure 400 invalid query parameter(s), see the Error field for details
// @Failure 500 failed to load property data
// @router / [get]
func (p *PropertyController) GetAllProperties() {

	// checking and parsing test on filters
	filter, errs := validators.ValidateAndParseFilter(&p.Controller)
	if errs != nil {
		utils.JsonError(&p.Controller, 400, errs)
		return
	}

	// take the property store
	store, err := services.GetPropertyStore()
	if err != nil {
		utils.JsonError(&p.Controller, 500, "failed to load property data")
		return
	}

	// find properties based on filter options
	result := services.GetFilteredProperties(store.GetData(), filter)
	utils.JsonSuccess(&p.Controller, 200, models.PropertyListResponse{Result: result})
}

// @Title Get A Property
// @Description Returns one rental property matching the given ID.
// @Param   id   path   string  true   "Property ID"
// @Success 200 {object} models.PropertyResponseDTO
// @Failure 400 id is not provided
// @Failure 404 property not found
// @router /:id [get]
func (p *PropertyController) GetAProperty() {

	// get id params from request
	id := p.Ctx.Input.Param(":id")
	if id == "" {
		utils.JsonError(&p.Controller, 400, "id is not provided")
		return
	}

	// get property store
	store, err := services.GetPropertyStore()
	if err != nil {
		utils.JsonError(&p.Controller, 500, "failed to load property data")
		return
	}

	// go to service layer and find property
	property, found := services.GetAPropertyByID(store.GetData(), id)
	if !found {
		utils.JsonError(&p.Controller, 404, "property not found")
		return
	}
	utils.JsonSuccess(&p.Controller, 200, *property)
}
