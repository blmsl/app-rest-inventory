package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	CateringTableName = "catering"
)

type Catering struct {
	Id       uint64    `orm:"auto"`
	Product  *Product  `orm:"rel(fk)"`
	Provider *Provider `orm:"rel(fk)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *Catering) TableName() string {
	return CateringTableName
}

type CateringDao struct {
	dao
}

func NewCateringDao(customerID string) *CateringDao {
	d := new(CateringDao)
	d.dao.CustomerID = customerID
	return d
}

func (d *CateringDao) FindByDates(start, end time.Time) ([]*Catering, error) {
	cond := orm.NewCondition().Or("Created__gte", start).Or("Created__lte", end)

	o := d.dao.getOrm()
	qs := o.QueryTable(CateringTableName).SetCond(cond)

	caterings := make([]*Catering, 0)
	err := readBy(qs, caterings)

	return caterings, err
}

func (d *CateringDao) StockValueByDates(start, end time.Time) (float64, error) {
	var value float64

	caterings, err := d.FindByDates(start, end)
	if err != nil {
		return value, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the stock value by dates.
	for _, catering := range caterings {
		value += catering.Product.Price
	}

	return value, nil
}
