package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	SaleTableName = "sale"
)

type Sale struct {
	Id          uint64       `orm:"auto"`
	Headquarter *Headquarter `orm:"rel(fk)"`
	Products    []*Product   `orm:"rel(m2m)"`
	UserId      string
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Discount    float64   `orm:"digits(10);decimals(2)"`
}

func (s *Sale) TableName() string {
	return SaleTableName
}

type SaleDao struct {
	dao
}

func NewSaleDao(customerID string) *SaleDao {
	d := new(SaleDao)
	d.dao.CustomerID = customerID
	return d
}

func (d *SaleDao) FindByDates(start, end time.Time) ([]*Sale, error) {
	cond := orm.NewCondition().Or("Created__gte", start).Or("Created__lte", end)

	o := d.dao.getOrm()
	qs := o.QueryTable(SaleTableName).SetCond(cond)

	sales := make([]*Sale, 0)
	err := readBy(qs, sales)

	return sales, err
}

func (d *SaleDao) RevenueByDates(start, end time.Time) (float64, error) {
	var revenue float64

	sales, err := d.FindByDates(start, end)
	if err != nil {
		return revenue, err
	}

	// Golang is faster than PostgreSQL SGBD so here we calc the revenue.
	for _, sale := range sales {
		// Add brute value.
		for _, product := range sale.Products {
			revenue += product.Price
		}

		// Add discount.
		if revenue != 0 && sale.Discount != 0 {
			revenue -= sale.Discount
		}
	}

	return revenue, nil
}
