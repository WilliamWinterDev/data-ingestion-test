package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WilliamWinterDev/test-metrics-api/metrics"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	host := os.Getenv("APP_HOST")

	if host == "" {
		host = ":6003"
	}

	router.HandleFunc("/api/metrics/{startTime}/{endTime}", metrics.Get).Methods("GET")
	router.HandleFunc("/api/metrics", metrics.Get).Methods("GET")

	fmt.Println("Listening on " + host)
	log.Fatal(http.ListenAndServe(host, router))
}
