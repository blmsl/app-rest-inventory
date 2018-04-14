package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	ProductTableName = "product"
)

type Product struct {
	Id      uint64 `orm:"auto"`
	Name    string
	Brand   string
	Color   string
	Amount  uint64
	Price   float64   `orm:"digits(10);decimals(2)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (p *Product) TableName() string {
	return ProductTableName
}

type ProductDao struct {
	dao
}

func NewProductDao(customerID string) *ProductDao {
	d := new(ProductDao)
	d.dao.CustomerID = customerID
	return d
}

func (d *ProductDao) FindByNameOrBrandOrColor(name, brand, color string) ([]*Product, error) {
	cond := orm.NewCondition().Or("Name__exact ", name).Or("Brand__exact", brand).Or("Color__exact", color)

	o := d.dao.getOrm()
	qs := o.QueryTable(ProductTableName).SetCond(cond)

	products := make([]*Product, 0)
	err := readBy(qs, products)

	return products, err
}

func (d *ProductDao) StockValue() (uint64, error) {
	var value uint64

	products := make([]*Product, 0)
	err := ReadAll(d.dao.CustomerID, products)
	if err != nil {
		return value, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock value.
	for _, product := range products {
		value += product.Amount
	}

	return value, nil

}
