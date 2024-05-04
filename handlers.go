package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


var people []Person

func peopleHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        getPeople(w, req)
    case http.MethodPost:
        postPeople(w, req)
    default:
        http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
    }
}

func getPeople(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
    fmt.Fprintf(w, "getPeople '%v'", people)
}

func postPeople(w http.ResponseWriter, req *http.Request) {
    var person Person
    err := json.NewDecoder(req.Body).Decode(&person)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    people = append(people, person)
    fmt.Fprintf(w, "post new person '%v'", person)
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "http web-server work correctly")
}
