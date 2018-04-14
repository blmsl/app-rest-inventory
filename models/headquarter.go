package models

import (
	"time"
)

type Headquarter struct {
	Id         uint64 `orm:"auto"`
	CustomerId string
	Name       string    `orm:"unique"`
	Address    string    `orm:"null;unique"`
	Phone      string    `orm:"null"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func (h *Headquarter) TableName() string {
	return "headquarter"
}
