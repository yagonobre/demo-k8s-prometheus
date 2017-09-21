package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func server() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/hostname", hostname)
	r.HandleFunc("/", hostnameWithDelay)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func main() {
	server()
}
