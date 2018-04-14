package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

// Caterings API
type CateringsController struct {
	beego.Controller
}

func (c *CateringsController) URLMapping() {
	c.Mapping("CreateCatering", c.CreateCatering)
}

// @Title CreateCatering
// @Description Create catering.
// @Accept json
// @router / [post]
func (c *CateringsController) CreateCatering() {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	catering := new(models.Catering)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, catering)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Insert catering.
	err = models.Insert(customerID, catering)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = catering
	c.ServeJSON()
}

// @Title GetCatering
// @Description Get catering.
// @Param	catering_id	path	uint64	true	"Catering id."
// @router /:catering_id [get]
func (c *CateringsController) GetCatering(catering_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate catering ID.
	if catering_id == nil {
		err := fmt.Errorf("catering_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	catering := new(models.Catering)
	catering.Id = *catering_id

	// Get the catering.
	err := models.Read(customerID, catering)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = catering
	c.ServeJSON()
}

// @Title GetCaterings
// @Description Get caterings.
// @Param from query time.Time false "From date"
// @Param to query time.Time false "To date"
// @router / [get]
func (c *CateringsController) GetCaterings(from, to time.Time) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewCateringDao(customerID)

	// Get caterings.
	caterings, err := dao.FindByDates(from, to)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Get stock value.
	value, _ := dao.StockValueByDates(from, to)

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(caterings)
	response["value"] = value
	response["caterings"] = caterings

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateCatering
// @Description Update catering.
// @Accept json
// @Param catering_id path uint64 true "Catering id."
// @router /:catering_id [patch]
func (c *CateringsController) UpdateCatering(catering_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate catering ID.
	if catering_id == nil {
		err := fmt.Errorf("catering_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	catering := new(models.Catering)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, catering)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	catering.Id = *catering_id

	// Update the catering.
	err = models.Update(customerID, catering)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = catering
	c.ServeJSON()
}

// @Title DeleteCatering
// @Description Delete catering.
// @Accept json
// @Param	catering_id	path	uint64	true	"Catering id."
// @router /:catering_id [delete]
func (c *ProductsController) DeleteCatering(catering_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate catering ID.
	if catering_id == nil {
		err := fmt.Errorf("catering_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	catering := new(models.Catering)
	catering.Id = *catering_id

	// Update the catering.
	err := models.Delete(customerID, catering)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}
