package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"time"
)

// Caterings API
type CateringsController struct {
	BaseController
}

func (c *CateringsController) URLMapping() {
	c.Mapping("CreateCatering", c.CreateCatering)
}

// @Title CreateCatering
// @Description Create catering.
// @Accept json
// @Success 200 {object} models.Catering
// @router / [post]
func (c *CateringsController) CreateCatering() {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Unmarshall request.
	catering := new(models.Catering)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, catering)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Insert catering.
	err = models.Insert(customerId, catering)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = catering
	c.ServeJSON()
}

// @Title GetCatering
// @Description Get catering.
// @Param	catering_id	path	uint64	true	"Catering id."
// @Success 200 {object} models.Catering
// @router /:catering_id [get]
func (c *CateringsController) GetCatering(catering_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Validate catering Id.
	if catering_id == nil {
		err := fmt.Errorf("catering_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	catering := new(models.Catering)
	catering.Id = *catering_id

	// Get the catering.
	err := models.Read(customerId, catering)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = catering
	c.ServeJSON()
}

// @Title GetCaterings
// @Description Get caterings.
// @Param from query time.Time false "From date"
// @Param to query time.Time false "To date"
// @Success 200 {object} map[string]interface{}
// @router / [get]
func (c *CateringsController) GetCaterings(from, to time.Time) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Build DAO.
	dao := models.NewCateringDao(customerId)

	// Get caterings.
	caterings, err := dao.FindByDates(from, to)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(caterings)
	response["caterings"] = caterings

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateCatering
// @Description Update catering.
// @Accept json
// @Param catering_id path uint64 true "Catering id."
// @Success 200 {object} models.Catering
// @router /:catering_id [patch]
func (c *CateringsController) UpdateCatering(catering_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Validate catering Id.
	if catering_id == nil {
		err := fmt.Errorf("catering_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Unmarshall request.
	catering := new(models.Catering)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, catering)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}
	catering.Id = *catering_id

	// Update the catering.
	err = models.Update(customerId, *catering_id, catering)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}

	// Serve JSON.
	c.Data["json"] = catering
	c.ServeJSON()
}

// @Title DeleteCatering
// @Description Delete catering.
// @Param	catering_id	path	uint64	true	"Catering id."
// @router /:catering_id [delete]
func (c *ProductsController) DeleteCatering(catering_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusUnauthorized, err.Error())
	}

	// Validate catering Id.
	if catering_id == nil {
		err := fmt.Errorf("catering_id can not be empty.")
		logs.Error(err.Error())
		c.serveError(http.StatusBadRequest, err.Error())
	}

	// Prepare query.
	catering := new(models.Catering)
	catering.Id = *catering_id

	// Update the catering.
	err := models.Delete(customerId, *catering_id, catering)
	if err != nil {
		logs.Error(err.Error())
		c.serveError(http.StatusInternalServerError, err.Error())
	}
}
