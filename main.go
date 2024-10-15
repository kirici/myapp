package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of GET requests by path.",
	},
	[]string{"code"},
)

func handleHello() http.HandlerFunc {
	return promhttp.InstrumentHandlerCounter(
		totalRequests,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello world!")
		}),
	)
}

func handleMath() http.HandlerFunc {
	return promhttp.InstrumentHandlerCounter(
		totalRequests,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, rand.Float64()*rand.Float64())
		}),
	)
}

func init() {
	prometheus.MustRegister(totalRequests) // Alternatively use .Register and check errs
}

func main() {
	http.Handle("/", handleHello())
	http.Handle("/math", handleMath())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
