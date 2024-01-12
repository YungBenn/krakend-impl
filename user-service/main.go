package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
    Text string `json:"text"`
}

func welcomeUserHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{"Welcome User!"}
    json.NewEncoder(w).Encode(msg)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{"Get User!"}
    json.NewEncoder(w).Encode(msg)
}

func main() {
    http.HandleFunc("/api", welcomeUserHandler)
    http.HandleFunc("/api/get", getUserHandler)
    http.ListenAndServe(":4001", nil)
}