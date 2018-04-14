package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Headquarters API
type HeadquartersController struct {
	beego.Controller
}

func (c *HeadquartersController) URLMapping() {
	c.Mapping("CreateHeadquarter", c.CreateHeadquarter)
	c.Mapping("GetHeadquarters", c.GetHeadquarters)
}

// @Title CreateHeadquarter
// @Description Create headquarter.
// @Accept json
// @router / [post]
func (c *HeadquartersController) CreateHeadquarter() {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	headquarter := new(models.Headquarter)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Insert headquarter.
	err = models.Insert(customerID, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = headquarter
	c.ServeJSON()
}

// @Title GetHeadquarter
// @Description Get headquarter.
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @router /:headquarter_id [get]
func (c *HeadquartersController) GetHeadquarter(headquarter_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter ID.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	headquarter := new(models.Headquarter)
	headquarter.Id = *headquarter_id

	// Get the headquarter.
	err := models.Read(customerID, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = headquarter
	c.ServeJSON()
}

// @Title GetHeadquarters
// @Description Get headquarters.
// @router / [get]
func (c *HeadquartersController) GetHeadquarters() {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	headquarters := make([]*models.Headquarter, 0)
	err := models.ReadAll(customerID, headquarters)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = headquarters
	c.ServeJSON()
}

// @Title UpdateHeadquarter
// @Description Update headquarter.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @router /:headquarter_id [patch]
func (c *HeadquartersController) UpdateHeadquarter(headquarter_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter ID.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	headquarter := new(models.Headquarter)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	headquarter.Id = *headquarter_id

	// Update the headquarter.
	err = models.Update(customerID, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = headquarter
	c.ServeJSON()
}

// @Title DeleteHeadquarter
// @Description Delete headquarter.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @router /:headquarter_id [delete]
func (c *HeadquartersController) DeleteHeadquarter(headquarter_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter ID.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	headquarter := new(models.Headquarter)
	headquarter.Id = *headquarter_id

	// Update the headquarter.
	err := models.Delete(customerID, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}
