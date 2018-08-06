package models

import (
	"errors"
		)

func setDefaultCheckRequired(required []string, defaultValues map[string]interface{}) (map[string]interface{}, error) {

	opt := map[string]interface{}{}

	for _, v := range required {
		_, ok := opt[v]
		if !ok {
			return opt, errors.New(v + " is required")
		}
	}

	for k, v := range defaultValues {
		_, ok := opt[k]
		if !ok {
			opt[k] = v
		}
	}

	return opt, nil
}

type ModelWithMethods interface {
	NewModel(map[string]interface{}) (interface{}, error)
	RetrieveOne() interface{}
	RetrieveArray() interface{}
}
