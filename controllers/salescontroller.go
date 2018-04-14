package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

// Sales API
type SalesController struct {
	beego.Controller
}

func (c *SalesController) URLMapping() {
	c.Mapping("CreateSale", c.CreateSale)
}

// @Title CreateSale
// @Description Create sale.
// @Accept json
// @router / [post]
func (c *SalesController) CreateSale() {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	sale := new(models.Sale)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, sale)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Insert sale.
	err = models.Insert(customerID, sale)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = sale
	c.ServeJSON()
}

// @Title GetSale
// @Description Get sale.
// @Param	sale_id	path	uint64	true	"Sale id."
// @router /:sale_id [get]
func (c *SalesController) GetSale(sale_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate sale ID.
	if sale_id == nil {
		err := fmt.Errorf("sale_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	sale := new(models.Sale)
	sale.Id = *sale_id

	// Get the sale.
	err := models.Read(customerID, sale)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = sale
	c.ServeJSON()
}

// @Title GetSales
// @Description Get sales.
// @Param from query time.Time false "From date"
// @Param to query time.Time false "To date"
// @router / [get]
func (c *SalesController) GetSales(from, to time.Time) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewSaleDao(customerID)

	// Get sales.
	sales, err := dao.FindByDates(from, to)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Get revenue.
	value, _ := dao.RevenueByDates(from, to)

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(sales)
	response["value"] = value
	response["sales"] = sales

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateSale
// @Description Update sale.
// @Accept json
// @Param sale_id path uint64 true "Sale id."
// @router /:sale_id [patch]
func (c *SalesController) UpdateSale(sale_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate sale ID.
	if sale_id == nil {
		err := fmt.Errorf("sale_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	sale := new(models.Sale)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, sale)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	sale.Id = *sale_id

	// Update the sale.
	err = models.Update(customerID, sale)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = sale
	c.ServeJSON()
}

// @Title DeleteSale
// @Description Delete sale.
// @Accept json
// @Param	sale_id	path	uint64	true	"Sale id."
// @router /:sale_id [delete]
func (c *ProductsController) DeleteSale(sale_id *uint64) {
	// Get customer ID from the cookies.
	customerID := c.Ctx.GetCookie("customer_id")
	if len(customerID) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate sale ID.
	if sale_id == nil {
		err := fmt.Errorf("sale_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	sale := new(models.Sale)
	sale.Id = *sale_id

	// Update the sale.
	err := models.Delete(customerID, sale)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}
