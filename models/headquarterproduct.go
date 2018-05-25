package models

import (
	"time"
)

var (
	HeadquarterProductTableName = "headquarter_product"
)

type HeadquarterProducts struct {
	Products []HeadquarterProduct `json:"products"`
}

type HeadquarterProduct struct {
	Id            uint64    `xorm:"pk autoincr"`
	HeadquarterId uint64    `xorm:"index"`
	ProductId     uint64    `xorm:"index"`
	Amount        uint64    `xorm:"not null"`
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
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

	// Get engine.
	engine := GetEngine(d.GetSchema())

	headquarterProducts := make([]*HeadquarterProductProduct, 0)
	err := engine.Table(HeadquarterProductTableName).Join("INNER", ProductTableName, "product.id = headquarter_product.product_id").
		Find(&headquarterProducts)
	if err != nil {
		return amount, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProduct := range headquarterProducts {
		amount += headquarterProduct.HeadquarterProduct.Amount
	}

	return amount, nil
}

// @Description Get the stock amount for specific headquarter.
// @Param headquarterId Headquarter Id.
func (d *HeadquarterProductDao) StockAmountByHeadquarter(headquarterId uint64) (uint64, error) {
	var amount uint64

	// Get engine.
	engine := GetEngine(d.GetSchema())

	headquarterProducts := make([]*HeadquarterProductProduct, 0)
	err := engine.Table(HeadquarterProductTableName).Join("INNER", ProductTableName, "product.id = headquarter_product.product_id").
		Where("headquarter_product.headquarter_id = ?", headquarterId).Find(&headquarterProducts)
	if err != nil {
		return amount, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProduct := range headquarterProducts {
		amount += headquarterProduct.HeadquarterProduct.Amount
	}

	return amount, nil
}

// @Description Get the stock price.
func (d *HeadquarterProductDao) StockPrice() (float64, error) {
	var total float64

	// Get engine.
	engine := GetEngine(d.GetSchema())

	headquarterProducts := make([]*HeadquarterProductProduct, 0)
	err := engine.Table(HeadquarterProductTableName).Join("INNER", ProductTableName, "product.id = headquarter_product.product_id").
		Find(&headquarterProducts)
	if err != nil {
		return total, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProduct := range headquarterProducts {
		total += float64(headquarterProduct.HeadquarterProduct.Amount) * headquarterProduct.Product.Price
	}

	return total, nil
}

// @Description Get the stock price for specific headquarter.
// @Param headquarterId Headquarter Id.
func (d *HeadquarterProductDao) StockPriceByHeadquarter(headquarterId uint64) (float64, error) {
	var total float64

	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	// TODO: Report a bug with JOIN to XORM.
	// In order to get the product details by headquarter, get the headquarter products and the get each product.
	// Get head quarters products.
	headquarterProducts := make([]*HeadquarterProduct, 0)
	err := engine.Where("headquarter_id = ?", headquarterId).Find(&headquarterProducts)
	if err != nil {
		return total, err
	}

	// Get the products and build result.
	for i := range headquarterProducts {
		headquarterProductProduct := new(HeadquarterProductProduct)

		// Setup headquarter.
		headquarterProduct := headquarterProducts[i]
		headquarterProductProduct.HeadquarterProduct = *headquarterProduct

		// Prepare query.
		product := new(Product)
		product.Id = headquarterProduct.ProductId
		// Get the product.
		err = Read(d.GetSchema(), product)
		if err != nil {
			return total, err
		}
		// Setup product.
		headquarterProductProduct.Product = *product

		headquarterProductProducts = append(headquarterProductProducts, headquarterProductProduct)
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock amount.
	for _, headquarterProductProduct := range headquarterProductProducts {
		total += float64(headquarterProductProduct.HeadquarterProduct.Amount) * headquarterProductProduct.Product.Price
	}

	return total, nil
}

// @Param headquarterId Headquarter Id.
// @Param name Product name.
// @Param brand Product brand.
// @Param color Product color.
func (d *HeadquarterProductDao) FindByHeadquarterOrNameOrBrandOrColor(headquarterId uint64, name, brand, color string) ([]*HeadquarterProductProduct, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())
	headquarterProductProducts := make([]*HeadquarterProductProduct, 0)

	// TODO: Report a bug with JOIN to XORM.
	// In order to get the product details by headquarter, get the headquarter products and the get each product.
	// Get head quarters products.
	headquarterProducts := make([]*HeadquarterProduct, 0)
	err := engine.Where("headquarter_id = ?", headquarterId).Find(&headquarterProducts)
	if err != nil {
		return nil, err
	}

	// Get the products and build result.
	for i := range headquarterProducts {
		headquarterProductProduct := new(HeadquarterProductProduct)

		// Setup headquarter.
		headquarterProduct := headquarterProducts[i]
		headquarterProductProduct.HeadquarterProduct = *headquarterProduct

		// Prepare query.
		product := new(Product)
		product.Id = headquarterProduct.ProductId
		// Get the product.
		err = Read(d.GetSchema(), product)
		if err != nil {
			return nil, err
		}

		headquarterProductProducts = append(headquarterProductProducts, headquarterProductProduct)

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
func (d *HeadquarterProductDao) Read(headquarterId, ProductId uint64) (*HeadquarterProductProduct, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	// Build Query.
	headquarterProduct := new(HeadquarterProductProduct)
	_, err := engine.Table(HeadquarterProductTableName).Join("INNER", ProductTableName, "product.id = headquarter_product.product_id").
		Where("headquarter_product.headquarter_id = ?", headquarterId).And("product.id = ?", ProductId).Get(&headquarterProduct)
	if err != nil {
		return nil, err
	}

	return headquarterProduct, err
}
