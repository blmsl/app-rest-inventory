package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// Register models here.
	orm.RegisterModel(
		new(Headquarter),
		new(Sale),
		new(Product),
		new(Catering),
		new(Provider))
}
