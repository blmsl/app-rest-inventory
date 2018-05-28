package models

import (
	"time"
)

var (
	ProviderTableName = "provider"
)

type Provider struct {
	Id      uint64    `xorm:"pk autoincr" json:"id"`
	Name    string    `xorm:"not null unique" json:"name"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

func (p *Provider) TableName() string {
	return ProviderTableName
}
