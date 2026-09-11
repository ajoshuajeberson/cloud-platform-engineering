package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Cloud Platform API is healthy",
		Status:  "UP",
	}

	json.NewEncoder(w).Encode(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Hello from the Cloud Platform!",
		Status:  "SUCCESS",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/hello", helloHandler)

	log.Println("Cloud Platform API starting on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
