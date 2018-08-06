package endpoints

import (
	"transations-calculator/db"
	"fmt"
	"net/http"
	"encoding/json"
)

type CategoriesEndpoint struct {
	Model db.ModelObject
}


func (c CategoriesEndpoint) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all")

	resArray := c.Model.Instance.RetrieveArray()

	errG := c.Model.GetAll(resArray, "")
	if errG != nil {
		SetResponseFail(errG.Error(), 400, &w)
		return
	}

	SetResponseSuccess(&w, resArray)
}

func (c CategoriesEndpoint) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one")
}

func (c CategoriesEndpoint) GetOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one")
}

func (c CategoriesEndpoint) UpdateOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one")
}

func (c CategoriesEndpoint) DeleteOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one")
}



func SetResponseFail(message string, code int, w *http.ResponseWriter) error {
	(*w).WriteHeader(code)
	return json.NewEncoder(*w).Encode(map[string]interface{}{"message": message, "code": code})
}

func SetResponseSuccess(w *http.ResponseWriter, response interface{}) {
	(*w).Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(*w).Encode(response)
}