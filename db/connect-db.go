package db

import (
	"gopkg.in/mgo.v2"
	"transations-calculator/models"
)

type Tables struct {
	Users ModelObject
}

var Categories ModelObject
var Transactions ModelObject

func GetTables(address string, dbName string) error {

	session, errDB := mgo.Dial(address)
	if errDB != nil {
		return errDB
	}

	db := session.DB(dbName)

	Categories = Model(db, "categories", models.Category{})
	Transactions = Model(db, "transactions", models.Transaction{})

	return nil
}
