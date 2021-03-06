package controllers

import (
	"app-rest-inventory/auth0"
	"app-rest-inventory/models"
	"app-rest-inventory/util/stringutil"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
)

const (
	Admin  = "admin"
	Seller = "seller"
)

// Customers API
type CustomersController struct {
	BaseController
}

func (c *CustomersController) URLMapping() {
	c.Mapping("CreateCustomer", c.CreateCustomer)
	c.Mapping("GetCustomer", c.GetCustomer)
}

// @Title CreateCustomer
// @Description Create customer.
// @Accept json
// @Success 200 {object} auth0.Group
// @router / [post]
func (c *CustomersController) CreateCustomer() {
	// Unmarshall request.
	customer := new(auth0.Group)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, customer)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Create customer.
	customer, err = auth0.Auth.CreateGroup(customer.Name, customer.Description)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Create the customer DB.
	go func(customerID string) {
		err = models.CreateCustomerSchema(customerID)
		if err != nil {
			logs.Error(err.Error())
		}
	}(customer.Id)

	// Create customer groups.
	go func(customerGroup *auth0.Group) {
		// Create nested groups.
		// Create admins group.
		nameBuilder := bytes.NewBufferString(customerGroup.Name)
		nameBuilder.WriteString(stringutil.HyphenMinus)
		nameBuilder.WriteString(Admin)
		var adminGroup *auth0.Group
		adminGroup, err = auth0.Auth.CreateGroup(nameBuilder.String(), fmt.Sprintf("%s admins group.", customerGroup.Name))
		if err != nil {
			logs.Error(err.Error())
		}

		// Create sellers group.
		nameBuilder.Reset()
		nameBuilder.WriteString(customerGroup.Name)
		nameBuilder.WriteString(stringutil.HyphenMinus)
		nameBuilder.WriteString(Seller)
		var sellerGroup *auth0.Group
		sellerGroup, err = auth0.Auth.CreateGroup(nameBuilder.String(), fmt.Sprintf("%s sellers group.", customerGroup.Name))
		if err != nil {
			logs.Error(err.Error())
		}

		// Nest groups.
		err = auth0.Auth.NestGroups(customerGroup.Id, adminGroup.Id, sellerGroup.Id)
		if err != nil {
			logs.Error(err.Error())
		}
	}(customer)

	// Serve JSON.
	c.Data["json"] = customer
	c.ServeJSON()
}

// @Title GetCustomer
// @Description Get customer.
// @Success 200 {object} auth0.Group
// @router / [get]
func (c *CustomersController) GetCustomer() {

	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	customer, err := auth0.Auth.GetGroup(customerID)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = customer
	c.ServeJSON()
}
