package main

import (
	"go-first/handlers"
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/people", handlers.PeopleHandler)
    http.HandleFunc("/health", handlers.HealthCheckHandler)

    err := http.ListenAndServe("localhost:8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
