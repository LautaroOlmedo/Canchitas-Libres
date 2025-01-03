package rest

import (
	"canchitas-libres/pkg/domain/person"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var persons []person.Person

type UserHandler struct {
}

func NewHandler() *UserHandler {
	return &UserHandler{}
}

func (handler *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost:
		handler.CreateUser(w, r)
		return
	case r.Method == http.MethodGet:
		handler.GetUsers(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}

func (handler *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json") // Seteo en los headers de la respuesta el formato que le quiero dar
		responseJson, err := json.Marshal(persons)         // Marshea "serealiza" a formato json
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(responseJson) // retorno la respuesta
	} else {
		w.Write([]byte("Not Found"))
	}

}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var person *person.Person
	w.Header().Set("Content-Type", "application/json") // Seteo en los headers de la respuesta el formato que le quiero dar
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		data, err := ioutil.ReadAll(r.Body) // lo que me mandan leo el body para convertirlo a un tipo de dato que entienda el compilador de Go
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(data, &person) // Si cuando envío hago Marshall, cuando recibo hago Unmarshall
		if err != nil {
			w.Write([]byte("invalid json"))
		}
		persons = append(persons, *person)

		w.WriteHeader(http.StatusCreated)
	}

}
