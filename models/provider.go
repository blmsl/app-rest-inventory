package models

import (
	"time"
)

var (
	ProviderTableName = "provider"
)

type Provider struct {
	Id      uint64 `xorm:"autoincr"`
	Name    string `xorm:"not null unique"`
	Address string
	Phone   string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func (p *Provider) TableName() string {
	return ProviderTableName
}
