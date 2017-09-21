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

func hostname(w http.ResponseWriter, r *http.Request) {
	hn, _ := os.Hostname()
	fmt.Fprintf(w, hn)
}

func hostnameWithDelay(w http.ResponseWriter, r *http.Request) {
	delay.Delay1()
	delay.Delay2()
	delay.Delay3()

	hn, _ := os.Hostname()
	fmt.Fprintf(w, hn)
}

func handle() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/hostname", hostname)
	r.HandleFunc("/", hostnameWithDelay)

	r.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handle()
}
