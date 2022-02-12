package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const (
	GET  string = "GET"
	POST string = "POST"
)

func main() {
	fmt.Println("Servidor")

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == GET {
			json.NewEncoder(w).Encode(" Funciona")
		}
		if r.Method == POST {
			json.NewEncoder(w).Encode(" Funciona POST")
		}

	})

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
