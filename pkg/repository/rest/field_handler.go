package rest

import (
	"canchitas-libres/pkg/domain/field"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var fields []field.Field

type FieldHandler struct {

}

func NewHandlerField() *FieldHandler {
	return &FieldHandler{}
}

func (h *FieldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch {
	case r.Method == http.MethodPost:
		h.CreateField(w, r)
		return
	case r.Method == http.MethodGet:
		h.GetFields(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}

func (h *FieldHandler) GetFields(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json") // Seteo en los headers de la respuesta el formato que le quiero dar
		responseJson, err := json.Marshal(fields)         // Marshea "serealiza" a formato json
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(responseJson) // retorno la respuesta
	} else {
		w.Write([]byte("Not Found"))
	}

}

func (h *FieldHandler) CreateField(w http.ResponseWriter, r *http.Request) {
	var field *field.Field
	w.Header().Set("Content-Type", "application/json") // Seteo en los headers de la respuesta el formato que le quiero dar
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		data, err := ioutil.ReadAll(r.Body) // lo que me mandan leo el body para convertirlo a un tipo de dato que entienda el compilador de Go
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.Unmarshal(data, &field) // Si cuando envío hago Marshall, cuando recibo hago Unmarshall
		if err != nil {
			w.Write([]byte("invalid json"))
		}
		fields = append(fields, *field)

		w.WriteHeader(http.StatusCreated)
	}

}