package endpoints

import (
	"transations-calculator/db"
	"fmt"
	"net/http"
)

type TransactionsEndpoint struct {
	Model db.ModelObject
}

func (c TransactionsEndpoint) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all")
}

func (c TransactionsEndpoint) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one")
}

func (c TransactionsEndpoint) GetOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one")
}

func (c TransactionsEndpoint) UpdateOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one")
}

func (c TransactionsEndpoint) DeleteOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one")
}
