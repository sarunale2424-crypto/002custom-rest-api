package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	//setup router

	router := http.NewServeMux()

	//setup routes

	router.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "GET method"})
	})

	router.HandleFunc("POST /items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "POST method"})
	})

	router.HandleFunc("PUT /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		id := r.PathValue("id")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "PUT method", "id": id})
	})

	router.HandleFunc("DELETE /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		id := r.PathValue("id")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "DELETE method", "id": id})
	})

	//setup server

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	slog.Info("Starting server at http://localhost:3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
