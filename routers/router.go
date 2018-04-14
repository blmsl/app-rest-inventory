package routers

import (
	"app-rest-inventory/controllers"
	"github.com/astaxie/beego"
)

// @APIVersion 1.0.0
// @Title GoStock API.
// @Description GoStock backend API.
// @Contact alobaton@gmail.com
// @TermsOfServiceUrl
// @License
// @LicenseUrl
func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/customers",
			beego.NSInclude(
				&controllers.CustomersController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
		beego.NSNamespace("/headquarters",
			beego.NSInclude(
				&controllers.HeadquartersController{},
			),
		),
		beego.NSNamespace("/sales",
			beego.NSInclude(
				&controllers.SalesController{},
			),
		),
		beego.NSNamespace("/products",
			beego.NSInclude(
				&controllers.ProductsController{},
			),
		),
		beego.NSNamespace("/caterings",
			beego.NSInclude(
				&controllers.CateringsController{},
			),
		),
		beego.NSNamespace("/providers",
			beego.NSInclude(
				&controllers.ProvidersController{},
			),
		),
	)
	// Register namespace.
	beego.AddNamespace(ns)
}
