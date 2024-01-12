package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
    Text string `json:"text"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{"Hello World!"}
    json.NewEncoder(w).Encode(msg)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{"Welcome Home!"}
    json.NewEncoder(w).Encode(msg)
}

func main() {
    http.HandleFunc("/api", homeHandler)
    http.HandleFunc("/api/world", helloWorldHandler)
    http.ListenAndServe(":4000", nil)
}