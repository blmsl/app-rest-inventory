package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"testing"
)

func init() {
	Driver = "postgres"

	Host = "ec2-54-243-213-188.compute-1.amazonaws.com"
	Port = "5432"
	Database = "d5o20r6igmuas6"

	User = "soqeiwwlwaxuex"
	Password = "e979d7be30202f18de7dd822758c98bc127bb5ba421a275ec968a6e9cccf768d"

	// Max open connections by default.
	MaxOpenConns = 20

	// Max cacher size.
	MaxCacherSize = 200

	// Build database connection chain.
	Chain = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", Host, Port, User, Password, Database)

	// Initialize the pool.
	pool = make(map[string]*xorm.Engine)
}

func TestCreateCustomerSchema(t *testing.T) {
	err := CreateCustomerSchema("customerid")
	if err != nil {
		t.Errorf("Error creating customer schema: %s.", err.Error())
	}
}
