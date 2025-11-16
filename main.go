package main

import (
	"002custom-rest-api/models"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"
)

var items []models.Item

func main() {
	//Initialise initial values for items
	items = []models.Item{
		{1, "Salt", 0.40, time.Now()},
		{2, "Rice", 5.40, time.Now()},
	}

	//setup router

	router := http.NewServeMux()

	//setup routes

	router.HandleFunc("GET /items", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "GET method", "items": items})
	})

	router.HandleFunc("POST /items", func(w http.ResponseWriter, r *http.Request) {
		var item models.Item
		json.NewDecoder(r.Body).Decode(&item)
		item.ID = len(items) + 1
		item.CreatedAt = time.Now()

		items = append(items, item)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"method_type": "POST method", "message": "successfully added item", "item": item})
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
