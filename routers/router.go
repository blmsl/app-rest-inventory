package routers

import (
	"app-rest-inventory/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"strings"
)

var CustomerIDFilter = func(ctx *context.Context) {
	// If business account creation.
	if strings.HasPrefix(ctx.Input.URL(), "/customers") && ctx.Input.IsPost() {
		return
	}
	// Get customer ID from the cookies.
	customerID := ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		fmt.Println("Paso por aquí")
		//c.serveError(http.StatusUnauthorized, err.Error())
	}
}

// @APIVersion 1.0.0
// @Title GoStock API.
// @Description GoStock backend API.
// @Contact alobaton@gmail.com
// @TermsOfServiceUrl
// @License
// @LicenseUrl
func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(CustomerIDFilter),
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
		beego.NSNamespace("/bills",
			beego.NSInclude(
				&controllers.BillsController{},
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
