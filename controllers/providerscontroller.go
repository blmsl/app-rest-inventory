package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Providers API
type ProvidersController struct {
	beego.Controller
}

func (c *ProvidersController) URLMapping() {
	c.Mapping("CreateProvider", c.CreateProvider)
	c.Mapping("GetProviders", c.GetProviders)
}

// @Title CreateProvider
// @Description Create provider.
// @Accept json
// @router / [post]
func (c *ProvidersController) CreateProvider() {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	provider := new(models.Provider)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, provider)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Insert provider.
	err = models.Insert(customerID, provider)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = provider
	c.ServeJSON()
}

// @Title GetProvider
// @Description Get provider.
// @Param	provider_id	path	uint64	true	"Provider id."
// @router /:provider_id [get]
func (c *ProvidersController) GetProvider(provider_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate provider ID.
	if provider_id == nil {
		err := fmt.Errorf("provider_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	provider := new(models.Provider)
	provider.Id = *provider_id

	// Get the provider.
	err := models.Read(customerID, provider)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = provider
	c.ServeJSON()
}

// @Title GetProviders
// @Description Get providers.
// @router / [get]
func (c *ProvidersController) GetProviders() {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	providers := make([]*models.Provider, 0)
	err := models.ReadAll(customerID, providers)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = providers
	c.ServeJSON()
}

// @Title UpdateProvider
// @Description Update provider.
// @Accept json
// @Param	provider_id	path	uint64	true	"Provider id."
// @router /:provider_id [patch]
func (c *ProvidersController) UpdateProvider(provider_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate provider ID.
	if provider_id == nil {
		err := fmt.Errorf("provider_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	provider := new(models.Provider)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, provider)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	provider.Id = *provider_id

	// Update the provider.
	err = models.Update(customerID, provider)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = provider
	c.ServeJSON()
}

// @Title DeleteProvider
// @Description Delete provider.
// @Accept json
// @Param	provider_id	path	uint64	true	"Provider id."
// @router /:provider_id [delete]
func (c *ProvidersController) DeleteProvider(provider_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate provider ID.
	if provider_id == nil {
		err := fmt.Errorf("provider_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	provider := new(models.Provider)
	provider.Id = *provider_id

	// Update the provider.
	err := models.Delete(customerID, provider)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}
