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

// get all properties controller, handles filter options
func (p *PropertyController) GetAllProperties() {

	// checking and parsing test on filters
	filter, errs := validators.ValidateAndParseFilter(&p.Controller)
	if errs != nil {
		utils.JsonError(&p.Controller, 400, map[string]interface{}{"Errors ": errs})
		return
	}

	// take the property store
	store, err := services.GetPropertyStore()
	if err != nil {
		utils.JsonError(&p.Controller, 500, map[string]interface{}{"Errors ": "failed to load property data"})
		return
	}

	// find properties based on filter options
	result := services.GetFilteredProperties(store.GetData(), filter)
	utils.JsonSuccess(&p.Controller, 200, models.PropertyListResponse{Result: result})
}

// get a single property controller without filter option
func (p *PropertyController) GetAProperty() {

	// get id params from request
	id := p.Ctx.Input.Param(":id")
	if id == "" {
		utils.JsonError(&p.Controller, 400, map[string]interface{}{"Errors :": "id is not provided"})
		return
	}

	// get property store
	store, err := services.GetPropertyStore()
	if err != nil {
		utils.JsonError(&p.Controller, 500, map[string]interface{}{"Errors ": "failed to load property data"})
		return
	}

	// go to service layer and find property
	property, found := services.GetAPropertyByID(store.GetData(), id)
	if !found {
		utils.JsonError(&p.Controller, 404, map[string]interface{}{"Errors ": "property not found"})
		return
	}
	utils.JsonSuccess(&p.Controller, 200, *property)
}
