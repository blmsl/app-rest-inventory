package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"reflect"
)

var (
	Driver       string
	Host         string
	User         string
	Password     string
	MaxOpenConns int
	MaxLifeTime  int
)

func init() {
	Driver = beego.AppConfig.String("database::driver")
	orm.RegisterDriver(Driver, orm.DRPostgres)

	Host = beego.AppConfig.String("database::host")
	User = beego.AppConfig.String("database::user")
	Password = beego.AppConfig.String("database::password")
	val, err := beego.AppConfig.Int("database::maxopenconns")
	if err != nil {
		logs.Error(err.Error())
	}
	MaxOpenConns = val
	val, err = beego.AppConfig.Int("database::maxlifetime")
	if err != nil {
		logs.Error(err.Error())
	}
	MaxLifeTime = val
}

//https://stackoverflow.com/questions/43377732/cannot-connect-to-postgresql-database-in-beego
//https://stackoverflow.com/questions/30235031/how-to-create-a-new-mysql-database-with-go-sql-driver
// @Param customerID Customer ID.
func CreateCustomerDB(customerID string) error {
	// Register de database.
	chain := fmt.Sprintf("%s:%s@%s/", User, Password, Host)
	orm.RegisterDataBase(customerID, Driver, chain)

	// Create the database.
	o := GetOrm(customerID)
	raw := o.Raw("CREATE DATABASE IF NOT EXISTS " + customerID)
	_, err := raw.Exec()
	if err != nil {
		return err
	}

	// Sync the database.
	err = orm.RunSyncdb(customerID, true, true)
	if err != nil {
		return err
	}

	return nil
}

// @Param customerID Customer ID.
func GetOrm(customerID string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(customerID)
	orm.SetMaxOpenConns(customerID, MaxOpenConns)
	return o
}

// @Param customerID Customer ID
// @Param model Model.
func Insert(customerID string, model interface{}) error {
	o := GetOrm(customerID)
	_, err := o.Insert(model)
	return err
}

// @Param customerID Customer ID
// @Param model Model.
func Read(customerID string, model interface{}) error {
	o := GetOrm(customerID)
	err := o.Read(model)
	return err
}

// @Param customerID Customer ID
// @Param models Models.
func ReadAll(customerID string, models interface{}) error {
	o := GetOrm(customerID)

	// Build query seter.
	t := reflect.TypeOf(models).Elem()
	model := reflect.New(t)
	qs := o.QueryTable(model)

	_, err := qs.All(models)
	return err
}

// @Param qs Quey seter.
// @Param models Models.
func readBy(qs orm.QuerySeter, models interface{}) error {
	_, err := qs.All(models)
	return err
}

// @Param customerID Customer ID
// @Param model Model.
func Update(customerID string, model interface{}) error {
	o := GetOrm(customerID)
	if o.Read(model) != nil {
		return fmt.Errorf("Not found.")
	}
	_, err := o.Update(model)
	return err
}

// @Param customerID Customer ID
// @Param model Model.
func Delete(customerID string, model interface{}) error {
	o := GetOrm(customerID)
	_, err := o.Delete(model)
	return err
}

type dao struct {
	CustomerID string
}

func (d *dao) getOrm() orm.Ormer {
	return GetOrm(d.CustomerID)
}
