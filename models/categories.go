package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Usage Example
// opt := map[string]string{
//     "name": "food",
// }
// category, err := NewCategory(opt)
//
// if err != nil {
//     fmt.Println(err)
// } else {
//     errInsert := Categories.Insert(category)
//     if errInsert != nil {
//         fmt.Println(errInsert)
//     }
// }

type Category struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CreatedDate time.Time     `bson:"name" json:"name"`
	ModelWithMethods
}

func (Category) NewModel(opt map[string]interface{}) (interface{}, error) {

	opt, errSet := setDefaultCheckRequired([]string{"name"}, map[string]interface{}{
		"createdDate": time.Now(),
	})
	if errSet != nil {
		return &Category{}, errSet
	}

	return &Category{
		Name:        opt["name"].(string),
		CreatedDate: opt["createdDate"].(time.Time),
	}, nil
}

func (Category) RetrieveOne() interface{} {
	return &Category{}
}

func (Category) RetrieveArray() interface{} {
	return &[]Category{}
}
