package models

import (
	"time"
)

var (
	ProductTableName = "product"
)

type Product struct {
	Id      uint64 `xorm:"autoincr"`
	Name    string `xorm:"not null unique"`
	Brand   string
	Color   string
	Price   float64   `xorm:"not null"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
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
