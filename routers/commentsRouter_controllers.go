package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomerController"] = append(beego.GlobalControllerRouter["app-rest-inventory/controllers:CustomerController"],
		beego.ControllerComments{
			Method: "CreateCustomer",
			Router: `/customer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
