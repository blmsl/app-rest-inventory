package models

import (
	"time"
)

var (
	CateringTableName = "catering"
)

type Catering struct {
	Id         uint64    `xorm:"pk autoincr" json:"id"`
	ProductId  uint64    `xorm:"index" json:"product_id"`
	ProviderId uint64    `xorm:"index" json:"provider_id"`
	Amount     uint64    `xorm:"not null" json:"amount"`
	Created    time.Time `xorm:"created" json:"created"`
	Updated    time.Time `xorm:"updated" json:"updated"`
}

func (c *Catering) TableName() string {
	return CateringTableName
}

// In order to access the product's price in a catering we need to
// do a join between catering, provider and product tables in the xorm way.
type CateringProviderProduct struct {
	Catering `xorm:"extends"`
	Provider `xorm:"extends"`
	Product  `xorm:"extends"`
}

type CateringDao struct {
	Dao
}

func NewCateringDao(schema string) *CateringDao {
	d := new(CateringDao)
	d.Dao = new(dao)
	d.SetSchema(schema)
	return d
}

// @Param start Start time.
// @Param end End time.
func (d *CateringDao) FindByDates(start, end time.Time) ([]*CateringProviderProduct, error) {
	// Get engine.
	engine := GetEngine(d.GetSchema())

	// Build Query.
	caterings := make([]*CateringProviderProduct, 0)
	err := engine.Table(CateringTableName).Join("INNER", ProductTableName, "product.id = catering.product_id").Join("INNER", ProviderTableName, "provider.id = catering.provider_id").
		Where("catering.created >= ?", start).Or("catering.created <= ?", end).Find(&caterings)

	return caterings, err
}
