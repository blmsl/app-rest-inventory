package models

import (
	"bytes"
	"fmt"
	"time"
)

var (
	HeadquarterProductTableName = "headquarter_product"
)

type HeadquarterProducts struct {
	Products []*HeadquarterProduct `json:"products"`
}

type HeadquarterProduct struct {
	Id            uint64    `xorm:"pk autoincr" json:"id"`
	HeadquarterId uint64    `xorm:"index" json:"headquarter_id"`
	ProductId     uint64    `xorm:"index" json:"product_id"`
	Amount        uint64    `xorm:"not null" json:"amount"`
	Created       time.Time `xorm:"created" json:"created"`
	Updated       time.Time `xorm:"updated" json:"updated"`
}

func (h *HeadquarterProduct) TableName() string {
	return HeadquarterProductTableName
}

// In order to access the information of the headquarter's products we need to
// do a join between headquarter_product and product in the xorm way.
type HeadquarterProductProduct struct {
	HeadquarterProduct `xorm:"extends"`
	Product            `xorm:"extends"`
}

type HeadquarterProductDao struct {
	Dao
}

// @Param schema Schema.
func NewHeadquarterProductDao(schema string) *HeadquarterProductDao {
	d := new(HeadquarterProductDao)
	d.Dao = new(dao)
	d.SetSchema(schema)
	return d
}

// @Description Get the stock amount.
func (d *HeadquarterProductDao) StockAmount() (uint64, error) {
	var amount uint64

	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(HeadquarterProductTableName)
	sql.WriteString(" hp INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON hp.product_id = p.id")

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	err := engine.Sql(sql.String()).Find(&headquarterProductProducts)
	if err != nil {
		return amount, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProductProduct := range headquarterProductProducts {
		amount += headquarterProductProduct.HeadquarterProduct.Amount
	}

	return amount, nil
}

// @Description Get the stock amount for specific headquarter.
// @Param headquarterId Headquarter Id.
func (d *HeadquarterProductDao) StockAmountByHeadquarter(headquarterId uint64) (uint64, error) {
	var amount uint64

	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(HeadquarterProductTableName)
	sql.WriteString(" hp INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON hp.product_id = p.id AND hp.headquarter_id = ")
	sql.WriteString(fmt.Sprintf("%v", headquarterId))

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	err := engine.Sql(sql.String()).Find(&headquarterProductProducts)
	if err != nil {
		return amount, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProductProduct := range headquarterProductProducts {
		amount += headquarterProductProduct.HeadquarterProduct.Amount
	}

	return amount, nil
}

// @Description Get the stock cost.
func (d *HeadquarterProductDao) StockCost() (float64, error) {
	var total float64

	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(HeadquarterProductTableName)
	sql.WriteString(" hp INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON hp.product_id = p.id")

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	// Execute the sentence.
	err := engine.Sql(sql.String()).Find(&headquarterProductProducts)
	if err != nil {
		return total, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProductProduct := range headquarterProductProducts {
		total += float64(headquarterProductProduct.HeadquarterProduct.Amount) * headquarterProductProduct.Product.Cost
	}

	return total, nil
}

// @Description Get the stock cost for specific headquarter.
// @Param headquarterId Headquarter Id.
func (d *HeadquarterProductDao) StockCostByHeadquarter(headquarterId uint64) (float64, error) {
	var total float64

	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(HeadquarterProductTableName)
	sql.WriteString(" hp INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON hp.product_id = p.id AND hp.headquarter_id = ")
	sql.WriteString(fmt.Sprintf("%v", headquarterId))

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	// Execute the sentence.
	err := engine.Sql(sql.String()).Find(&headquarterProductProducts)
	if err != nil {
		return total, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProductProduct := range headquarterProductProducts {
		total += float64(headquarterProductProduct.HeadquarterProduct.Amount) * headquarterProductProduct.Product.Cost
	}

	return total, nil
}

// @Param headquarterId Headquarter Id.
// @Param name Product name.
// @Param brand Product brand.
// @Param color Product color.
func (d *HeadquarterProductDao) FindByHeadquarterOrNameOrBrandOrColor(headquarterId uint64, name, brand, color string) ([]*HeadquarterProductProduct, error) {
	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(HeadquarterProductTableName)
	sql.WriteString(" hp INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON hp.product_id = p.id AND hp.headquarter_id = ")
	sql.WriteString(fmt.Sprintf("%v", headquarterId))
	if len(name) > 0 || len(brand) > 0 || len(color) > 0 {
		sql.WriteString(" WHERE ")
	}
	if len(name) > 0 {
		sql.WriteString("p.name LIKE '")
		sql.WriteString(name)
		sql.WriteString("%'")
	}
	if len(name) > 0 && len(brand) > 0 {
		sql.WriteString(" OR ")
	}
	if len(brand) > 0 {
		sql.WriteString("p.brand = '")
		sql.WriteString(brand)
		sql.WriteString("'")
	}
	if (len(name) > 0 || len(brand) > 0) && len(color) > 0 {
		sql.WriteString(" OR ")
	}
	if len(color) > 0 {
		sql.WriteString("p.color = '")
		sql.WriteString(color)
		sql.WriteString("'")
	}

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	// Execute sentence.
	err := engine.Sql(sql.String()).AllCols().Find(&headquarterProductProducts)
	if err != nil {
		return nil, err
	}

	return headquarterProductProducts, err
}

// @Param headquarterId Headquarter Id.
// @Param productId Product Id.
func (d *HeadquarterProductDao) DeleteByHeadquarterIdAndProductId(headquarterId, productId uint64) error {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	headquarterProduct := new(HeadquarterProduct)
	headquarterProduct.HeadquarterId = headquarterId
	headquarterProduct.ProductId = productId

	_, err := engine.Delete(&headquarterProduct)

	return err
}

// @Param headquarterId Headquarter Id.
// @Param productId Product Id.
func (d *HeadquarterProductDao) Read(headquarterId, productId uint64) (*HeadquarterProductProduct, error) {
	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT * FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(HeadquarterProductTableName)
	sql.WriteString(" hp INNER JOIN ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)
	sql.WriteString(" p ON hp.product_id = p.id AND hp.headquarter_id = ")
	sql.WriteString(fmt.Sprintf("%v", headquarterId))
	sql.WriteString(" AND p.id = ")
	sql.WriteString(fmt.Sprintf("%v", productId))

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	// Execute the sentence.
	err := engine.Sql(sql.String()).Find(&headquarterProductProducts)
	if err != nil {
		return nil, err
	}

	length := len(headquarterProductProducts)
	if length != 1 {
		err = fmt.Errorf("headquarterProductProducts does not have the appropriate length. Current %d, Expected %d", length, 1)
		return nil, err
	}

	return headquarterProductProducts[0], nil
}

func (d *HeadquarterProductDao) Update(headquarterId, productId uint64, product *HeadquarterProduct) error {
	// Get engine.
	engine := GetEngine(d.GetSchema())
	_, err := engine.Update(product, &HeadquarterProduct{HeadquarterId: headquarterId, ProductId: productId})
	if err != nil {
		return err
	}

	return nil
}
