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
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
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
	err = models.Insert(customerId, headquarter)
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
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter Id.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	headquarter := new(models.Headquarter)
	headquarter.Id = *headquarter_id

	// Get the headquarter.
	err := models.Read(customerId, headquarter)
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
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	headquarters := make([]*models.Headquarter, 0)
	err := models.ReadAll(customerId, &headquarters)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(headquarters)
	response["headquarters"] = headquarters

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title UpdateHeadquarter
// @Description Update headquarter.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @router /:headquarter_id [patch]
func (c *HeadquartersController) UpdateHeadquarter(headquarter_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter Id.
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
	err = models.Update(customerId, *headquarter_id, headquarter)
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
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter Id.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	headquarter := new(models.Headquarter)
	headquarter.Id = *headquarter_id

	// Update the headquarter.
	err := models.Delete(customerId, *headquarter_id, headquarter)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}

// @Title AddProducts
// @Description Add products to headquarter.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @router /:headquarter_id/products [patch]
func (c *HeadquartersController) AddProducts(headquarter_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter Id.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	productsId := make([]uint64, 0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, productsId)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Create inserts.
	var errors []error
	for _, productId := range productsId {
		headquarterProduct := new(models.HeadquarterProduct)
		headquarterProduct.HeadquarterId = *headquarter_id
		headquarterProduct.ProductId = productId
		headquarterProduct.Amount = 0

		errors = append(errors, models.Insert(customerId, headquarterProduct))
	}

	if len(errors) > 0 {
		err := fmt.Errorf("Errors adding products.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}

// @Title GetProduct
// @Description Get headquarter product.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @Param	product_id	path	uint64	true	"Product id."
// @router /:headquarter_id/products/:product_id [get]
func (c *HeadquartersController) GetProduct(headquarter_id, product_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	// Validate headquarter Id.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate product Id.
	if product_id == nil {
		err := fmt.Errorf("product_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewHeadquarterProductDao(customerId)

	// Get the product.
	product, err := dao.Read(*headquarter_id, *product_id)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = product.Product
	c.ServeJSON()
}

// @Title GetProduct
// @Description Get headquarter product.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @Param name query string false "Product name."
// @Param brand query string false "Product brand."
// @Param color query string false "Product color."
// @router /:headquarter_id/products [get]
func (c *HeadquartersController) GetProducts(headquarter_id *uint64, name, brand, color string) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter Id.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewHeadquarterProductDao(customerId)

	// Get products.
	products, err := dao.FindByHeadquarterOrNameOrBrandOrColor(*headquarter_id, name, brand, color)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	ps := make([]models.Product, 0)
	for _, product := range products {
		ps = append(ps, product.Product)
	}

	// Get stock price.
	price, err := dao.StockPriceByHeadquarter(*headquarter_id)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["total"] = len(ps)
	response["price"] = price
	response["products"] = ps

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title RemoveProducts
// @Description Remove products from headquarter.
// @Accept json
// @Param	headquarter_id	path	uint64	true	"Headquarter id."
// @router /:headquarter_id/products [delete]
func (c *HeadquartersController) RemoveProducts(headquarter_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate headquarter Id.
	if headquarter_id == nil {
		err := fmt.Errorf("headquarter can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	productsId := make([]uint64, 0)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, productsId)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build dao.
	dao := models.NewHeadquarterProductDao(customerId)

	// Create inserts.
	var errors []error
	for _, productId := range productsId {

		errors = append(errors, dao.DeleteByHeadquarterIdAndProductId(*headquarter_id, productId))
	}

	if len(errors) > 0 {
		err := fmt.Errorf("Errors removing products.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}
