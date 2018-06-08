package models

import (
	"app-rest-inventory/util/stringutil"
	"bytes"
	"fmt"
	"time"
)

var (
	SaleTableName = "sale"
)

// @Description Sale or bill item.
type Sale struct {
	Id        uint64    `xorm:"pk autoincr" json:"id"`
	BillId    uint64    `xorm:"index" json:"bill_id"`
	ProductId uint64    `xorm:"index" json:"product_id"`
	Amount    uint64    `xorm:"not null" json:"amount"`
	Created   time.Time `xorm:"created" json:"created"`
	Updated   time.Time `xorm:"updated" json:"updated"`
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

// FindByBill
// @Param BillId Bill Id.
func (d *SaleDao) FindByBill(billId uint64) ([]*SaleBillProduct, error) {
	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(SaleTableName)
	sql.WriteString(" s INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(BillTableName)
	sql.WriteString(" b ON s.bill_id = b.id ")
	sql.WriteString("INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" b ON s.product_id = p.id ")
	sql.WriteString("WHERE s.bill_id = ")
	sql.WriteString(fmt.Sprintf("%v", billId))
	sql.WriteString(" ORDER BY s.id ASC")

	// Get engine.
	engine := GetEngine(d.GetSchema())
	sales := make([]*SaleBillProduct, 0)

	// Execute sentence.
	err := engine.Sql(sql.String()).AllCols().Find(&sales)
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
	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(SaleTableName)
	sql.WriteString(" s INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(BillTableName)
	sql.WriteString(" b ON s.bill_id = b.id ")
	sql.WriteString("INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON s.product_id = p.id ")
	sql.WriteString("AND s.product_id = ")
	sql.WriteString(fmt.Sprintf("%v", productId))
	sql.WriteString(" WHERE s.created >= '")
	sql.WriteString(start.UTC().Format(stringutil.UTCFormat))
	sql.WriteString("' OR s.created <= '")
	sql.WriteString(end.UTC().Format(stringutil.UTCFormat))
	sql.WriteString("' ORDER BY s.id ASC")

	// Get engine.
	engine := GetEngine(d.GetSchema())
	sales := make([]*SaleBillProduct, 0)

	// Execute sentence.
	err := engine.Sql(sql.String()).AllCols().Find(&sales)
	if err != nil {
		return nil, err
	}

	return sales, err
}

// @Description Get sales by dates.
// @Param start Start time.
// @Param end End time.
func (d *SaleDao) FindByDates(start, end time.Time) ([]*SaleBillProduct, error) {
	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(SaleTableName)
	sql.WriteString(" s INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(BillTableName)
	sql.WriteString(" b ON s.bill_id = b.id ")
	sql.WriteString("INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON s.product_id = p.id ")
	sql.WriteString("WHERE s.created >= '")
	sql.WriteString(start.UTC().Format(stringutil.UTCFormat))
	sql.WriteString("' OR s.created <= '")
	sql.WriteString(end.UTC().Format(stringutil.UTCFormat))
	sql.WriteString("' ORDER BY s.id ASC")

	// Get engine.
	engine := GetEngine(d.GetSchema())
	sales := make([]*SaleBillProduct, 0)

	// Execute sentence.
	err := engine.Sql(sql.String()).AllCols().Find(&sales)
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
