package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dimiro1/graphql-api/products"
	"github.com/dimiro1/graphql-api/resolvers"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
)

func main() {
	file, err := os.Open("schema.graphql")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	db := sqlx.MustOpen("sqlite3", ":memory:")

	db.MustExec(`CREATE TABLE Products (
                        id INTEGER PRIMARY KEY, 
                        name TEXT,
                        price REAL
                );
                INSERT INTO Products VALUES (1, 'TV', 500);
                INSERT INTO Products VALUES (2, 'Microwave', 100);
                INSERT INTO Products VALUES (3, 'Refrigerator', 800);
                INSERT INTO Products VALUES (4, 'Dishwasher', 400);`)

	inventoryRepository := &products.DatabaseInventoryRepository{DB: db}

	schema := graphql.MustParseSchema(string(data), &resolvers.QueriesResolver{
		InventoryRepository: inventoryRepository,
	})

	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Print("Starting to listen 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
