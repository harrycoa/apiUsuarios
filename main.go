package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/harrycoa/apiUsuarios/handlers"
	"net/http"
	"time"
)

func main() {
	handlerGeneral := handlers.NewHandlerGeneral()

	fmt.Println("Servidor")

	router := mux.NewRouter()

	// Rutas
	router.HandleFunc("/", handlerGeneral.HandleGeneral)

	router.HandleFunc("/{id}", handlerGeneral.HandleOne)

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
