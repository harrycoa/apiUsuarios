package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/harrycoa/apiUsuarios/entities"
	"github.com/harrycoa/apiUsuarios/services"
	"net/http"
)

const (
	GET  string = "GET"
	POST string = "POST"
	PUT string = "PUT"
	DELETE string = "DELETE"
)

type UsuarioController interface {
	HandleGeneral(w http.ResponseWriter, r *http.Request)
	HandleOne(w http.ResponseWriter, r *http.Request)
}

// constructor
type usuarioController struct {
	service services.UsuarioService
}

func NewUsuarioController() UsuarioController {
	return &usuarioController{
		service: services.NewUsuarioService(),

	}
}

// Attach
func (h *usuarioController) HandleGeneral(w http.ResponseWriter, r *http.Request) {

	if r.Method == GET {
		usuarios := h.service.FindAll()
		json.NewEncoder(w).Encode(" Funciona")
	}
	if r.Method == POST {
		usuario := entities.Usuario{}
		json.NewDecoder(r.Body).Decode(&usuario)
		h.service.Save()
		json.NewEncoder(w).Encode(" Funciona POST")
	}
	if r.Method == PUT {
		id, encontrado := mux.Vars(r)["id"]
		if encontrado {
			usuario := entities.Usuario{}
			json.NewDecoder(r.Body).Decode(&usuario)
			usuarioNew := h.service.Update(id, usuario)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(usuarioNew)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Error encontrado")
	}

}

func (h *usuarioController) HandleOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == POST {
		id, encontrado := mux.Vars(r)["id"]

		if encontrado {
			usuario,err := h.service.FindOne(id)

			if err != nil{
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(usuario)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(nil)
	}
}
