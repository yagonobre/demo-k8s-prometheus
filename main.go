package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "demo_requests_total",
		Help: "The current number of requests.",
	})
)

func init() {
	prometheus.MustRegister(requestsTotal)
}

func server() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/hostname", hostname)
	r.HandleFunc("/", hostnameWithDelay)

	r.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
		ConnState: func(c net.Conn, s http.ConnState) {
			switch s {
			case http.StateNew:
				requestsTotal.Inc()
			case http.StateHijacked | http.StateClosed:
				requestsTotal.Dec()
			}
		},
	}

	log.Fatal(s.ListenAndServe())
}

func main() {
	server()
}
