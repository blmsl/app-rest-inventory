package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"],
		beego.ControllerComments{
			Method: "CreateCustomer",
			Router: `/customers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/users`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
