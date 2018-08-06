package main

import (
	"transations-calculator/endpoints"
	"transations-calculator/config"
	"transations-calculator/db"
	"fmt"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()

	errGetTables := db.GetTables(*config.DB_ADDR, *config.DB_NAME)
	if errGetTables != nil {
		fmt.Errorf("failed to get tables: %v", errGetTables)
	}
	fmt.Println("Setup DB connection")

	if *config.IsSeed {
		seed()
	} else {
		premain()
	}
}

func seed() {
	db.Categories.Delete(nil)
	fmt.Println("Deleted all user")

	db.Transactions.Delete(nil)
	fmt.Println("Removed all invites")
}

func premain() {
	c := endpoints.CategoriesEndpoint{Model: db.Categories}
	t := endpoints.TransactionsEndpoint{Model: db.Transactions}

	m := mux.NewRouter().StrictSlash(true)

	endpoints.SetCrud(m, c, "category")
	endpoints.SetCrud(m, t, "transaction")

	fmt.Errorf(http.ListenAndServe(*config.HttpAddr, m).Error())
}
