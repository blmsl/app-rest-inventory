package models

import (
	"time"
)

var (
	HeadquarterTableName = "headquarter"
)

type Headquarter struct {
	Id         uint64    `xorm:"pk autoincr" json:"id"`
	CustomerId string    `xorm:"index" json:"customer_id"`
	Name       string    `xorm:"not null unique" json:"name"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	Created    time.Time `xorm:"created" json:"created"`
	Updated    time.Time `xorm:"updated" json:"updated"`
}

func (h *Headquarter) TableName() string {
	return HeadquarterTableName
}
