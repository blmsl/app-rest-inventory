package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Products API
type ProductsController struct {
	beego.Controller
}

func (c *ProductsController) URLMapping() {
	c.Mapping("CreateProduct", c.CreateProduct)
	c.Mapping("GetBrands", c.GetBrands)
}

// @Title CreateProduct
// @Description Create product.
// @Accept json
// @Success 200 {object} models.Product
// @Failure 403 :product_id is empty
// @router / [post]
func (c *ProductsController) CreateProduct() {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	product := new(models.Product)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, product)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Insert product.
	err = models.Insert(customerId, product)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = product
	c.ServeJSON()
}

// @Title GetProduct
// @Description Get product.
// @Param	product_id	path	uint64	true	"Product id."
// @Success 200 {object} models.Product
// @Failure 403 :product_id is empty
// @router /:product_id [get]
func (c *ProductsController) GetProduct(product_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate product Id.
	if product_id == nil {
		err := fmt.Errorf("product_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	product := new(models.Product)
	product.Id = *product_id

	// Get the product.
	err := models.Read(customerId, product)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = product
	c.ServeJSON()
}

// @Title GetProducts
// @Description Get products.
// @Param name query string false "Product name."
// @Param brand query string false "Product brand."
// @Param color query string false "Product color."
// @router / [get]
func (c *ProductsController) GetProducts(name, brand, color string) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewProductDao(customerId)

	// Get products.
	products, err := dao.FindByNameOrBrandOrColor(name, brand, color)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(products)
	response["products"] = products

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateProduct
// @Description Update product.
// @Accept json
// @Param	product_id	path	uint64	true	"Product id."
// @Success 200 {object} models.Product
// @Failure 403 :product_id is empty
// @router /:product_id [patch]
func (c *ProductsController) UpdateProduct(product_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate product Id.
	if product_id == nil {
		err := fmt.Errorf("product_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	product := new(models.Product)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, product)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	product.Id = *product_id

	// Update the product.
	err = models.Update(customerId, *product_id, product)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = product
	c.ServeJSON()
}

// @Title DeleteProduct
// @Description Delete product.
// @Accept json
// @Param	product_id	path	uint64	true	"Product id."
// @Failure 403 :product_id is empty
// @router /:product_id [delete]
func (c *ProductsController) DeleteProduct(product_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate product Id.
	if product_id == nil {
		err := fmt.Errorf("product_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	product := new(models.Product)
	product.Id = *product_id

	// Update the product.
	err := models.Delete(customerId, *product_id, product)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}

// @Title GetBrands
// @Description Get product brands.
// @Accept json
// @router /brands [get]
func (c *ProductsController) GetBrands() {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewProductDao(customerId)

	// Get brands.
	brands, err := dao.GetBrands()
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(brands)
	response["brands"] = brands

	c.Data["json"] = response
	c.ServeJSON()

}
