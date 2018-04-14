package models

import (
	"time"
)

type Provider struct {
	Id      uint64    `orm:"auto"`
	Name    string    `orm:"unique"`
	Address string    `orm:"null"`
	Phone   string    `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (p *Provider) TableName() string {
	return "provider"
}
