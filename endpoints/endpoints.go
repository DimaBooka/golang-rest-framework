package endpoints

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Endpoints interface {
	GetAll(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	GetOne(http.ResponseWriter, *http.Request)
	UpdateOne(http.ResponseWriter, *http.Request)
	DeleteOne(http.ResponseWriter, *http.Request)
}

func SetCrud(m *mux.Router, endpoints Endpoints, pass string) {
	baseRoute := "/" + pass
	withUuid := baseRoute + "/{uuid}"

	m.HandleFunc(baseRoute, endpoints.GetAll).Methods("GET")
	m.HandleFunc(baseRoute, endpoints.Post).Methods("POST")
	m.HandleFunc(withUuid, endpoints.GetOne).Methods("GET")
	m.HandleFunc(withUuid, endpoints.UpdateOne).Methods("PUT")
	m.HandleFunc(withUuid, endpoints.DeleteOne).Methods("DELETE")
}
