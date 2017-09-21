package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/yagonobre/demo-k8s-prometheus/delay"
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
