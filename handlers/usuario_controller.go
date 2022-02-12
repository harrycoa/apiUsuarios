package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	GET  string = "GET"
	POST string = "POST"
)

type InterfaceGeneral interface {
	HandleGeneral(w http.ResponseWriter, r *http.Request)
	HandleOne(w http.ResponseWriter, r *http.Request)
}

// constructor
type HandlerGeneral struct {
}

func NewHandlerGeneral() InterfaceGeneral {
	return HandlerGeneral{}
}

// Attach
func (h HandlerGeneral) HandleGeneral(w http.ResponseWriter, r *http.Request) {

	if r.Method == GET {
		json.NewEncoder(w).Encode(" Funciona")
	}
	if r.Method == POST {
		json.NewEncoder(w).Encode(" Funciona POST")
	}

}

func (h HandlerGeneral) HandleOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		id, encontrado := mux.Vars(r)["id"]

		if encontrado {
			cadena := "se recibio el id = " + id
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(cadena)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}
