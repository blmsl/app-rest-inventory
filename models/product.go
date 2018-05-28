package models

import (
	"bytes"
	"time"
)

var (
	ProductTableName = "product"
)

type Product struct {
	Id      uint64    `xorm:"pk autoincr" json:"id"`
	Name    string    `xorm:"not null unique" json:"name"`
	Brand   string    `json:"brand"`
	Color   string    `json:"color"`
	Price   float64   `xorm:"not null" json:"price"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

func (p *Product) TableName() string {
	return ProductTableName
}

type ProductDao struct {
	Dao
}

func NewProductDao(schema string) *ProductDao {
	d := new(ProductDao)
	d.Dao = new(dao)
	d.SetSchema(schema)
	return d
}

// @Param name Product name.
// @Param brand Product brand.
// @Param color Product color.
func (d *ProductDao) FindByNameOrBrandOrColor(name, brand, color string) ([]*Product, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	// Build Query.
	products := make([]*Product, 0)
	err := engine.Where("name = ?", name).Or("brand = ?", brand).Or("color = ?", color).Find(&products)

	return products, err
}

func (d *ProductDao) GetBrands() ([]string, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	// Build sentence.
	var sql bytes.Buffer
	sql.WriteString("SELECT DISTINCT(brand) FROM ")
	sql.WriteString("\"")
	sql.WriteString(d.GetSchema())
	sql.WriteString("\".")
	sql.WriteString(ProductTableName)

	// Execute sentence.
	brands := make([]string, 0)
	err := engine.Sql(sql.String()).Find(&brands)
	if err != nil {
		return nil, err
	}

	return brands, nil
}
