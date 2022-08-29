package db

import (
	"database/sql"
	"github.com/RossiniM/full-cycle-hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var db *sql.DB

func setUp() {
	db, _ = sql.Open("sqlite3", ":memory:")
	createTable(db)
	createProducts(db)
}

func createTable(db2 *sql.DB) {
	table := ` create table products(id string, name string, price float,status string);`
	statement, err := db2.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func createProducts(db2 *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, "disabled")`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setUp()
	defer db.Close()
	productDB := New(db)
	productInterface, err := productDB.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "disabled", productInterface.GetStatus())
	require.Equal(t, 0.0, productInterface.GetPrice())
	require.Equal(t, "Product Test", productInterface.GetName())

}

func TestProductDB_Save(t *testing.T) {
	setUp()
	defer db.Close()
	productDB := New(db)
	name := "Product"
	price := 50.0
	product := application.NewProduct(&name, &price)

	productSaved, err := productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, productSaved.GetName())
	require.Equal(t, product.Price, productSaved.GetPrice())
	require.Equal(t, product.Status, productSaved.GetStatus())

	product.Price = 0.0
	product.Status = "disabled"

	productUpdated, err := productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Price, productUpdated.GetPrice())
	require.Equal(t, product.Status, productUpdated.GetStatus())

}
