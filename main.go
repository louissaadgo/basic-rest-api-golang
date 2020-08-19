package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type article struct {
	Title string
	Body  string
}

func first(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Try localhost:8080/api")
}

func api(w http.ResponseWriter, r *http.Request) {

	data := []article{
		article{Title: "First Title", Body: "First Body"},
		article{Title: "SEC Title", Body: "SEC Body"},
	}
	json.NewEncoder(w).Encode(data)
}

func handleRequest() {
	http.HandleFunc("/", first)
	http.HandleFunc("/api", api)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
