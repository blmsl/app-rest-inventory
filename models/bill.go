package models

import (
	"time"
)

var (
	BillTableName = "bill"
)

type Bill struct {
	Id            uint64    `xorm:"pk autoincr"`
	HeadquarterId uint64    `xorm:"index"`
	UserId        string    `xorm:"index"`
	Discount      float64   `xorm:"not null"`
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
}

func (b *Bill) TableName() string {
	return BillTableName
}
