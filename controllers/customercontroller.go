package controllers

import (
	"app-rest-inventory/auth0"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Customer API
type CustomerController struct {
	beego.Controller
}

func (c *CustomerController) URLMapping() {
	c.Mapping("CreateCustomer", c.CreateCustomer)
}

// @router /customer [post]
func (c *CustomerController) CreateCustomer() {
	// Unmarshall request.
	g := new(auth0.Group)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, g)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Create customer.
	g, err = auth0.AUTH0.CreateGroup(g.Name, g.Description)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = g
	c.ServeJSON()
}
