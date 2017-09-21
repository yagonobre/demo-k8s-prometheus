package main

import (
	"fmt"
	"net/http"
	"os"

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
