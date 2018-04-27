package models

import (
	"time"
)

var (
	SaleTableName = "sale"
)

// @Description Sale or bill item.
type Sale struct {
	Id        uint64    `xorm:"autoincr"`
	BillId    uint64    `xorm:"index"`
	ProductId uint64    `xorm:"index"`
	Amount    uint64    `xorm:"not null"`
	Created   time.Time `xorm:"created"`
	Updated   time.Time `xorm:"updated"`
}

func (s *Sale) TableName() string {
	return SaleTableName
}

// In order to access the product's sales in a bill we need to
// do a join between bill, sale and product tables in the xorm way.
type SaleBillProduct struct {
	Sale    `xorm:"extends"`
	Bill    `xorm:"extends"`
	Product `xorm:"extends"`
}

type SaleDao struct {
	Dao
}

func NewSaleDao(schema string) *SaleDao {
	d := new(SaleDao)
	d.Dao = new(dao)
	d.SetSchema(schema)
	return d
}

// @Param BillId Bill Id.
func (d *SaleDao) FindByBill(BillId uint64) ([]*SaleBillProduct, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	sales := make([]*SaleBillProduct, 0)
	err := engine.Table(SaleTableName).Join("INNER", BillTableName, "bill.id = sale.bill_id").Join("INNER", ProductTableName, "product.id = sale.product_id").
		Where("sale.bill_id = ?", BillId).Asc("sale.id").Find(&sales)
	if err != nil {
		return nil, err
	}

	return sales, err
}

// @Description Get product sales by dates.
// @Param productId Product Id.
// @Param start Start time.
// @Param end End time.
func (d *SaleDao) FindByProductAndDates(productId uint64, start, end time.Time) ([]*SaleBillProduct, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	// Build Query.
	sales := make([]*SaleBillProduct, 0)
	err := engine.Table(SaleTableName).Join("INNER", BillTableName, "bill.id = sale.bill_id").Join("INNER", ProductTableName, "product.id = sale.product_id").
		Where("sale.product_id = ?", productId).And("sale.created >= ?", start).Or("sale.created <= ?", end).Asc("sale.bill_id").Find(&sales)
	if err != nil {
		return nil, err
	}

	return sales, err
}

// @Description Get sales by dates.
// @Param start Start time.
// @Param end End time.
func (d *SaleDao) FindByDates(start, end time.Time) ([]*SaleBillProduct, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	// Build Query.
	sales := make([]*SaleBillProduct, 0)
	err := engine.Table(SaleTableName).Join("INNER", BillTableName, "bill.id = sale.bill_id").Join("INNER", ProductTableName, "product.id = sale.product_id").
		Where("sale.created >= ?", start).Or("sale.created <= ?", end).Asc("sale.bill_id").Find(&sales)
	if err != nil {
		return nil, err
	}

	return sales, err
}

// @Description Get revenue by dates.
// @Param start Start time.
// @Param end End time.
func (d *SaleDao) RevenueByDates(start, end time.Time) (float64, error) {
	var revenue float64

	// Find by dates and then group by bill programatically.
	sales, err := d.FindByDates(start, end)
	if err != nil {
		return revenue, err
	}
	// Group by bill.
	bills := make(map[uint64][]*SaleBillProduct)
	for _, sale := range sales {
		bills[sale.Sale.BillId] = append(bills[sale.Sale.BillId], sale)
	}

	// Calculate revenue.
	for _, bSales := range bills {
		// Calculate bill revenue.
		var billRevenue float64
		for _, sale := range bSales {
			billRevenue += float64(sale.Sale.Amount) * sale.Product.Price
		}
		// Apply discount.
		if len(bSales) > 0 {
			billRevenue -= bSales[0].Bill.Discount
		}

		// Add bill revenue.
		revenue += billRevenue
	}

	return revenue, nil
}
