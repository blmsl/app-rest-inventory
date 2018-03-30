package controllers

import (
	"app-rest-inventory/auth0"
	"app-rest-inventory/util/stringutil"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	Admin  = "admin"
	Seller = "seller"
)

// Customer API
type CustomersController struct {
	beego.Controller
}

func (c *CustomersController) URLMapping() {
	c.Mapping("CreateCustomer", c.CreateCustomer)
}

// @router /customers [post]
func (c *CustomersController) CreateCustomer() {
	// Unmarshall request.
	customerGroup := new(auth0.Group)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, customerGroup)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Create customer.
	customerGroup, err = auth0.AUTH0.CreateGroup(customerGroup.Name, customerGroup.Description)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	go func(customerGroup *auth0.Group) {
		// Create nested groups.
		// Create admins group.
		nameBuilder := bytes.NewBufferString(customerGroup.Name)
		nameBuilder.WriteString(stringutil.HyphenMinus)
		nameBuilder.WriteString(Admin)
		var adminGroup *auth0.Group
		adminGroup, err = auth0.AUTH0.CreateGroup(nameBuilder.String(), fmt.Sprintf("%s admins group.", customerGroup.Name))
		if err != nil {
			logs.Error(err.Error())
			c.Abort(err.Error())
		}

		// Create sellers group.
		nameBuilder.Reset()
		nameBuilder.WriteString(customerGroup.Name)
		nameBuilder.WriteString(stringutil.HyphenMinus)
		nameBuilder.WriteString(Seller)
		var sellerGroup *auth0.Group
		sellerGroup, err = auth0.AUTH0.CreateGroup(nameBuilder.String(), fmt.Sprintf("%s sellers group.", customerGroup.Name))
		if err != nil {
			logs.Error(err.Error())
			c.Abort(err.Error())
		}

		// Nest groups.
		err = auth0.AUTH0.NestGroups(customerGroup.Id, adminGroup.Id, sellerGroup.Id)
		if err != nil {
			logs.Error(err.Error())
			c.Abort(err.Error())
		}
	}(customerGroup)

	// Serve JSON.
	c.Data["json"] = customerGroup
	c.ServeJSON()
}
