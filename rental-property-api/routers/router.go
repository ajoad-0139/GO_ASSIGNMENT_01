package routers

import (
	"rental-property-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// versioned routing
	v1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/properties",
			beego.NSRouter("/", &controllers.PropertyController{}, "get:GetAllProperties"),
			beego.NSRouter("/:id", &controllers.PropertyController{}, "get:GetAProperty"),
		),
	)
	beego.AddNamespace(v1)
}
