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

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomersController"],
		beego.ControllerComments{
			Method: "GetCustomer",
			Router: `/customers/:customer_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("customer_id", param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/users`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetUsers",
			Router: `/users`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/users/:user_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(
				param.New("user_id", param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:UsersController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/users/:user_id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(
				param.New("user_id", param.InPath),
			),
			Params: nil})

}
