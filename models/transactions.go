package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

// Usage Example
// opt := map[string]string{
//     "name": "food",
// }
// transaction, err := NewTransaction(opt)
//
// if err != nil {
//     fmt.Println(err)
// } else {
//     errInsert := Transactions.Insert(transaction)
//     if errInsert != nil {
//         fmt.Println(errInsert)
//     }
// }

type Transaction struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CreatedDate time.Time     `bson:"name" json:"name"`
	ModelWithMethods
}

func (Transaction) NewModel(opt map[string]interface{}) (interface{}, error) {

	opt, errSet := setDefaultCheckRequired([]string{"name"}, map[string]interface{}{
		"createdDate": time.Now(),
	})
	if errSet != nil {
		return &Transaction{}, errSet
	}

	return &Transaction{
		Name:        opt["name"].(string),
		CreatedDate: opt["createdDate"].(time.Time),
	}, nil
}

func (Transaction) RetrieveOne() interface{}{
	return &Transaction{}
}

func (Transaction) RetrieveArray() interface{}{
	return &[]Transaction{}
}