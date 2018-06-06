package controllers

import (
	"app-rest-inventory/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type Bill struct {
	Id            uint64  `json:"id"`
	HeadquarterId uint64  `json:"headquarter_id"`
	UserId        string  `json:"user_id"`
	Discount      float64 `json:"discount"`
	Sales         []*Sale `json:"sales"`
}

type Sale struct {
	Id      uint64   `json:"id"`
	Amount  uint64   `json:"amount"`
	Product *Product `json:"product"`
}

type Product struct {
	Id    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Bills API
type BillsController struct {
	beego.Controller
}

func (c *BillsController) URLMapping() {
	c.Mapping("CreateBill", c.CreateBill)
}

// @Title CreateBill
// @Description Create bill.
// @Accept json
// @router / [post]
func (c *BillsController) CreateBill() {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	request := new(Bill)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, request)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Insert bill.
	b := new(models.Bill)
	b.HeadquarterId = request.HeadquarterId
	b.UserId = request.UserId
	b.Discount = request.Discount
	err = models.Insert(customerId, b)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
	request.Id = b.Id

	// Insert sales.
	var errors []error
	for _, sale := range request.Sales {
		s := new(models.Sale)
		s.BillId = b.Id
		s.ProductId = sale.Product.Id
		s.Amount = sale.Amount

		err = models.Insert(customerId, s)
		errors = append(errors, err)

		sale.Id = s.Id
	}

	if len(errors) > 0 {
		err := fmt.Errorf("Errors creating sales.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	c.Data["json"] = request
	c.ServeJSON()
}

// @Title GetBill
// @Description Get bill.
// @Param	bill_id	path	uint64	true	"Bill id."
// @router /:bill_id [get]
func (c *BillsController) GetBill(bill_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate sale Id.
	if bill_id == nil {
		err := fmt.Errorf("bill_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	bill := new(models.Bill)
	bill.Id = *bill_id

	// Get the bill.
	err := models.Read(customerId, bill)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Get the sales.
	dao := models.NewSaleDao(customerId)
	sales, err := dao.FindByBill(*bill_id)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build response.
	response := new(Bill)
	response.Id = bill.Id
	response.HeadquarterId = bill.HeadquarterId
	response.UserId = bill.UserId
	response.Discount = bill.Discount
	response.Sales = make([]*Sale, 0)

	for _, sale := range sales {
		s := new(Sale)
		s.Id = sale.Sale.Id
		s.Amount = sale.Sale.Amount
		s.Product = new(Product)
		s.Product.Id = sale.Sale.ProductId
		s.Product.Name = sale.Product.Name
		s.Product.Price = sale.Product.Price

		response.Sales = append(response.Sales, s)
	}

	// Serve JSON.
	c.Data["json"] = response
	c.ServeJSON()
}

// @Title GetBills
// @Description Get bills.
// @Param from query time.Time false "From date"
// @Param to query time.Time false "To date"
// @router / [get]
func (c *BillsController) GetBills(from, to time.Time) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Build DAO.
	dao := models.NewSaleDao(customerId)

	// Get sales.
	sales, err := dao.FindByDates(from, to)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Group by bill.
	bills := make(map[uint64][]*models.SaleBillProduct)
	for _, sale := range sales {
		bills[sale.Sale.BillId] = append(bills[sale.Sale.BillId], sale)
	}

	// Build response.
	response := make([]*Bill, 0)
	for BillId, sales := range bills {
		// Build bill.
		b := new(Bill)
		b.Id = BillId
		if len(sales) > 0 {
			b.HeadquarterId = sales[0].Bill.HeadquarterId
			b.UserId = sales[0].Bill.UserId
			b.Discount = sales[0].Bill.Discount
			b.Sales = make([]*Sale, 0)
		}

		for _, sale := range sales {
			s := new(Sale)
			s.Id = sale.Sale.Id
			s.Amount = sale.Sale.Amount
			s.Product = new(Product)
			s.Product.Id = sale.Sale.ProductId
			s.Product.Name = sale.Product.Name
			s.Product.Price = sale.Product.Price

			b.Sales = append(b.Sales, s)
		}

		// Add to response.
		response = append(response, b)
	}

	// Get revenue.
	revenue, _ := dao.RevenueByDates(from, to)

	// Serve JSON.
	rs := make(map[string]interface{})
	rs["total"] = len(response)
	rs["revenue"] = revenue
	rs["bills"] = response

	c.Data["json"] = rs
	c.ServeJSON()
}

// @Title UpdateDiscount
// @Description Update bill discount.
// @Accept json
// @Param bill_id path uint64 true "Bill id."
// @router /:bill_id [patch]
func (c *BillsController) UpdateDiscount(bill_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate sale Id.
	if bill_id == nil {
		err := fmt.Errorf("bill_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Unmarshall request.
	request := new(Bill)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, request)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Update the bill.
	b := new(models.Bill)
	b.Id = *bill_id
	b.Discount = request.Discount
	err = models.Update(customerId, *bill_id, b)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Serve JSON.
	response := make(map[string]interface{})
	response["discount"] = b.Discount

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title AddSale
// @Description Add sale to bill.
// @Accept json
// @Param bill_id path uint64 true "Bill id."
// @Param sale_id path uint64 true "Sale id."
// @router /:bill_id/sales/:sale_id [patch]
func (c *BillsController) AddSale(bill_id, sale_id *uint64) {
	// TODO
}

// @Title RemoveSale
// @Description Remove sale from bill.
// @Accept json
// @Param bill_id path uint64 true "Bill id."
// @Param sale_id path uint64 true "Sale id."
// @router /:bill_id/sales/:sale_id [delete]
func (c *BillsController) RemoveSale(bill_id, sale_id *uint64) {
	// TODO
}

// @Title DeleteBill
// @Description Delete bill.
// @Accept json
// @Param	bill_id	path	uint64	true	"Bill id."
// @router /:bill_id [delete]
func (c *BillsController) DeleteBill(bill_id *uint64) {
	// Get customer Id from the cookies.
	customerId := c.Ctx.GetCookie("customer_id")
	if len(customerId) == 0 {
		err := fmt.Errorf("customer_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Validate bill Id.
	if bill_id == nil {
		err := fmt.Errorf("bill_id can not be empty.")
		logs.Error(err.Error())
		c.Abort(err.Error())
	}

	// Prepare query.
	b := new(models.Bill)
	b.Id = *bill_id

	// Update the sale.
	err := models.Delete(customerId, *bill_id, b)
	if err != nil {
		logs.Error(err.Error())
		c.Abort(err.Error())
	}
}
