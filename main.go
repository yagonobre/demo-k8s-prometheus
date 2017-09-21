package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func server() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/hostname", hostname)
	r.HandleFunc("/", hostnameWithDelay)

	r.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func main() {
	server()
}
