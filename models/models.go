package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	Driver        string
	Host          string
	Port          string
	Database      string
	User          string
	Password      string
	MaxOpenConns  int
	MaxCacherSize int
	Chain         string

	// We need to maintain a pool of connection pools configured for each client
	// and allow specialized resource allocation. pool is a pool of *xorm.Engine.
	ExpirationTime  int
	CleanupInterval int
	pool            *cache.Cache
)

// Init models package.
func init() {
	Driver = beego.AppConfig.String("database::driver")

	Host = beego.AppConfig.String("database::host")
	Port = beego.AppConfig.String("database::port")
	Database = beego.AppConfig.String("database::database")

	User = beego.AppConfig.String("database::user")
	Password = beego.AppConfig.String("database::password")

	// Max open connections by default.
	val, err := beego.AppConfig.Int("database::maxopenconns")
	if err != nil {
		logs.Error(err.Error())
	}
	MaxOpenConns = val

	// Max cacher size.
	val, err = beego.AppConfig.Int("database::maxcachersize")
	if err != nil {
		logs.Error(err.Error())
	}
	MaxCacherSize = val

	// Build database connection chain.
	Chain = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", Host, Port, User, Password, Database)

	// Initialize the pool.
	val, err = beego.AppConfig.Int("database::expirationtime")
	if err != nil {
		logs.Error(err.Error())
	}
	ExpirationTime = val

	val, err = beego.AppConfig.Int("database::cleanupinterval")
	if err != nil {
		logs.Error(err.Error())
	}
	CleanupInterval = val

	pool = cache.New(time.Duration(ExpirationTime)*time.Minute, time.Duration(CleanupInterval)*time.Minute)
}

// @Param customerID Customer ID.
func CreateCustomerSchema(customerID string) error {
	/** customerID = strings.ToLower(customerID) */
	// Create new engine.
	engine, err := xorm.NewEngine(Driver, Chain)
	if err != nil {
		return err
	}

	// Create the schema.
	sql := fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS "%s" AUTHORIZATION %s`, customerID, User)
	/*logs.Debug(sql)*/
	_, err = engine.Exec(sql)
	if err != nil {
		return err
	}

	// Setup the search path to the engine.
	engine.SetSchema(customerID)

	// Setup cache to perform querys.
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), MaxCacherSize)
	engine.SetDefaultCacher(cacher)

	// Setup location.
	engine.TZLocation = time.Now().Location()

	// Setup max open connections.
	engine.SetMaxOpenConns(MaxOpenConns)

	// Add the engine to the pool.
	pool.Set(customerID, engine, time.Duration(ExpirationTime)*time.Minute)

	// Sync the tables.
	err = engine.Sync2(new(Bill), new(Catering), new(Headquarter), new(HeadquarterProduct), new(Product), new(Provider), new(Sale))
	if err != nil {
		logs.Error(err.Error())
		return err
	}

	return nil
}

// @Param customerID Customer ID.
func GetEngine(customerID string) *xorm.Engine {
	/** customerID = strings.ToLower(customerID) */
	// Validate engine.
	e, found := pool.Get(customerID)
	if !found {
		// Create new engine.
		engine, err := xorm.NewEngine(Driver, Chain)
		if err != nil {
			logs.Error(err.Error())
			return nil
		}

		// Setup the search path to the engine.
		engine.SetSchema(customerID)

		// Setup cache to perform querys.
		cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), MaxCacherSize)
		engine.SetDefaultCacher(cacher)

		// Setup location.
		engine.TZLocation = time.Now().Location()

		// Setup max open connections.
		engine.SetMaxOpenConns(MaxOpenConns)

		// Sync the tables.
		err = engine.Sync2(new(Bill), new(Catering), new(Headquarter), new(HeadquarterProduct), new(Product), new(Provider), new(Sale))
		if err != nil {
			logs.Error(err.Error())
			return nil
		}

		// Add the engine to the pool.
		pool.Set(customerID, engine, time.Duration(ExpirationTime)*time.Minute)
	}
	return e.(*xorm.Engine)
}

// @Param customerID Customer ID
// @Param model Model.
func Insert(customerID string, model interface{}) error {
	engine := GetEngine(customerID)
	_, err := engine.Insert(model)
	return err
}

// @Param customerID Customer ID
// @Param model Model.
func Read(customerID string, model interface{}) error {
	engine := GetEngine(customerID)
	_, err := engine.Get(model)
	return err
}

// @Param customerID Customer ID
// @Param models Models.
func ReadAll(customerID string, models interface{}) error {
	engine := GetEngine(customerID)
	err := engine.Find(models)
	return err
}

// @Param customerID Customer ID
// @Param modelID Model ID.
// @Param model Model.
func Update(customerID string, modelID interface{}, model interface{}) error {
	engine := GetEngine(customerID)

	_, err := engine.ID(modelID).Update(model)
	return err
}

// @Param customerID Customer ID
// @Param modelID Model ID.
// @Param model Model.
func Delete(customerID string, modelID interface{}, model interface{}) error {
	engine := GetEngine(customerID)

	_, err := engine.ID(modelID).Delete(model)
	return err
}

// Dao interface.
type Dao interface {
	GetSchema() string
	SetSchema(string)
}

// Dao struct definition.
type dao struct {
	Schema string
}

// Return the customer schema.
func (d *dao) GetSchema() string {
	return d.Schema
}

// @Param schema The customer schema.
func (d *dao) SetSchema(schema string) {
	d.Schema = schema
}
