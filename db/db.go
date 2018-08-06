package db

import (
	"reflect"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"transations-calculator/models"
)

type ModelObject struct {
	collection *mgo.Collection
	Instance models.ModelWithMethods
}

func Model(db *mgo.Database, collectionName string, model models.ModelWithMethods) ModelObject {
	c := db.C(collectionName)
	m := ModelObject{collection: c, Instance: model}
	return m
}

func (m *ModelObject) ShowColName() string {
	return m.collection.Name
}

func (m *ModelObject) GetAllCount() int {
	n, _ := m.collection.Count()
	return n
}

func (m *ModelObject) Insert(data interface{}) error {
	return m.collection.Insert(&data)
}

func (m *ModelObject) Update(update interface{}, selector interface{}) (info *mgo.ChangeInfo, err error) {
	if selector == nil {
		selector = bson.D{}
	}

	updateData := bson.M{"$set": update}
	return m.collection.UpdateAll(selector, updateData)
}

func (m *ModelObject) Upsert(update interface{}, selector interface{}) (info *mgo.ChangeInfo, err error) {
	return m.collection.Upsert(selector, update)
}

func (m *ModelObject) Delete(selector interface{}) (info *mgo.ChangeInfo, err error) {
	if selector == nil {
		selector = bson.D{}
	}
	return m.collection.RemoveAll(selector)
}

func (m *ModelObject) DeleteOne(selector interface{}) error {
	return m.collection.Remove(selector)
}

func (m *ModelObject) GetOne(selector interface{}, result interface{}, sortField string) (err error) {
	if selector == nil {
		selector = bson.D{}
	}

	if len(sortField) == 0 {
		sortField = "id"
	}

	v := reflect.ValueOf(result)
	switch v.Kind() {
	case reflect.Ptr:
		break
	case reflect.Struct:
		return status.Error(codes.Aborted, "modelObject can't deal with struct values, use a pointer")
	}

	err = m.collection.Find(selector).Sort(sortField).One(result)

	if err != nil {
		return err
	}
	return nil
}

func (m *ModelObject) Get(selector interface{}, resArray interface{}, sortField string) (err error) {
	if selector == nil {
		selector = bson.D{}
	}

	if len(sortField) == 0 {
		sortField = "id"
	}

	v := reflect.ValueOf(resArray)
	switch v.Kind() {
	case reflect.Ptr:
		break
	case reflect.Struct:
		return status.Error(codes.Aborted, "modelObject can't deal with struct values, use a pointer")
	default:
		return status.Error(codes.Aborted, "modelObject needs a map or a pointer to a struct")
	}

	err = m.collection.Find(selector).Sort(sortField).All(resArray)

	if err != nil {
		return err
	}
	return nil
}

func (m *ModelObject) GetAll(resArray interface{}, sortField string) error {
	return m.Get(nil, resArray, sortField)
}

func (m *ModelObject) Distinct(resArray interface{}, field string) error {
	return m.collection.Find(nil).Distinct(field, resArray)
}
