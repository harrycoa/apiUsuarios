package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"www.github.com/harrycoa/apiUsuarios/handlers"
)

func main() {
	usuarioController := handlers.NewUsuarioController()

	fmt.Println("Servidor")

	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/", usuarioController.HandleGeneral)

	router.HandleFunc("/{id}", usuarioController.HandleOne)

	// Servidor
	server := http.Server{

		Addr:              ":4000",
		Handler:           router,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	server.ListenAndServe()

}
