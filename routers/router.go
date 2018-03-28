package routers

import (
	"app-rest-inventory/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.CustomerController{})
}