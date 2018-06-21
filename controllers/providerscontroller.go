package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
)

// Providers API
type ProvidersController struct {
	BaseController
}

func (c *ProvidersController) URLMapping() {
	c.Mapping("CreateProvider", c.CreateProvider)
	c.Mapping("GetProviders", c.GetProviders)
}

// @Title CreateProvider
// @Description Create provider.
// @Accept json
// @Success 200 {object} models.Provider
// @router / [post]
func (c *ProvidersController) CreateProvider() {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Unmarshall request.
	provider := new(models.Provider)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, provider)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Insert provider.
	err = models.Insert(customerId, provider)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = provider
	c.ServeJSON()
}

// @Title GetProvider
// @Description Get provider.
// @Param	provider_id	path	uint64	true	"Provider id."
// @Success 200 {object} models.Provider
// @router /:provider_id [get]
func (c *ProvidersController) GetProvider(provider_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Validate provider Id.
	if provider_id == nil {
		err := fmt.Errorf("provider_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	provider := new(models.Provider)
	provider.Id = *provider_id

	// Get the provider.
	err := models.Read(customerId, provider)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = provider
	c.ServeJSON()
}

// @Title GetProviders
// @Description Get providers.
// @Success 200 {object} map[string]interface{}
// @router / [get]
func (c *ProvidersController) GetProviders() {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	providers := make([]*models.Provider, 0)
	err := models.ReadAll(customerId, providers)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(providers)
	response["providers"] = providers

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateProvider
// @Description Update provider.
// @Accept json
// @Param	provider_id	path	uint64	true	"Provider id."
// @Success 200 {object} models.Provider
// @router /:provider_id [patch]
func (c *ProvidersController) UpdateProvider(provider_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Validate provider Id.
	if provider_id == nil {
		err := fmt.Errorf("provider_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	provider := new(models.Provider)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, provider)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}
	provider.Id = *provider_id

	// Update the provider.
	err = models.Update(customerId, *provider_id, provider)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = provider
	c.ServeJSON()
}

// @Title DeleteProvider
// @Description Delete provider.
// @Param	provider_id	path	uint64	true	"Provider id."
// @router /:provider_id [delete]
func (c *ProvidersController) DeleteProvider(provider_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Validate provider Id.
	if provider_id == nil {
		err := fmt.Errorf("provider_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	provider := new(models.Provider)
	provider.Id = *provider_id

	// Update the provider.
	err := models.Delete(customerId, *provider_id, provider)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}
}
