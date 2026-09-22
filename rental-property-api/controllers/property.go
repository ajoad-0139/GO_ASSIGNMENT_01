package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

// get all properties controller, handles filter options
func (p *PropertyController) GetAllProperties() {

}

// get a single property controller without filter option
func (p *PropertyController) GetAProperty() {

}
