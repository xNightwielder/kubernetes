package main

import (
	"fmt"

	"net/http"

	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		opsProcessed.Inc()
		fmt.Fprintf(w, "Uygulama Çalışıyor!\n")
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		http.Get("/healthz")
	})

	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
