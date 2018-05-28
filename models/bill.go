package models

import (
	"time"
)

var (
	BillTableName = "bill"
)

type Bill struct {
	Id            uint64    `xorm:"pk autoincr" json:"id"`
	HeadquarterId uint64    `xorm:"index" json:"headquarter_id"`
	UserId        string    `xorm:"index" json:"user_id"`
	Discount      float64   `xorm:"not null" json:"discount"`
	Created       time.Time `xorm:"created" json:"created"`
	Updated       time.Time `xorm:"updated" json:"updated"`
}

func (b *Bill) TableName() string {
	return BillTableName
}
