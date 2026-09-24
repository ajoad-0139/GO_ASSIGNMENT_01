// @APIVersion 1.0.0
// @Title Rental Property API
// @Description REST API for browsing and filtering rental properties.
// @Contact support@example.com
// @TermsOfServiceUrl http://swagger.io/terms/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rental-property-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	//versioned routing
	v1 := beego.NewNamespace("/v1",
		beego.NSNamespace("/properties",
			beego.NSRouter("/", &controllers.PropertyController{}, "get:GetAllProperties"),
			beego.NSRouter("/:id", &controllers.PropertyController{}, "get:GetAProperty"),
		),
	)
	beego.AddNamespace(v1)
}
