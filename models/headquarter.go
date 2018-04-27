package models

import (
	"time"
)

var (
	HeadquarterTableName = "headquarter"
)

type Headquarter struct {
	Id         uint64 `xorm:"autoincr"`
	CustomerId string `xorm:"index"`
	Name       string `xorm:"not null unique"`
	Address    string
	Phone      string
	Created    time.Time `xorm:"created"`
	Updated    time.Time `xorm:"updated"`
}

func (h *Headquarter) TableName() string {
	return HeadquarterTableName
}
