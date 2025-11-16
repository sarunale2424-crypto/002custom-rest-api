package main

import (
	"002custom-rest-api/models"
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

var items []models.Item

func findItem(id int) error {
	for index := range items {
		if items[index].ID == id {
			return nil
		}
	}
	return errors.New("specified id not found")

}
func main() {
	//Initialise initial values for items
	items = []models.Item{
		{ID: 1, Name: "Salt", Price: 0.40, CreatedAt: time.Now()},
		{ID: 2, Name: "Rice", Price: 5.40, CreatedAt: time.Now()},
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
		var item models.Item

		var newItems []models.Item

		idInt, err := strconv.Atoi(id)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "invlalid ID," + err.Error()})
			return
		}

		err = findItem(idInt)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"method": "PUT method", "error": err.Error()})
			return

		} else {
			for index := range items {
				if items[index].ID != idInt {
					newItems = append(newItems, items[index])
				} else {
					json.NewDecoder(r.Body).Decode(&item)
					item.ID = idInt
					item.CreatedAt = time.Now()
					newItems = append(newItems, item)
				}
			}
			items = newItems

		}

		json.NewEncoder(w).Encode(map[string]interface{}{"method": "PUT method", "item": item, "message": "successfully updated item"})

	})

	router.HandleFunc("DELETE /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		id := r.PathValue("id")

		var newItems []models.Item

		idInt, err := strconv.Atoi(id)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "invlalid ID," + err.Error()})
			return
		}

		err = findItem(idInt)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"method": "PUT method", "error": err.Error()})
			return

		} else {
			for index := range items {
				if items[index].ID != idInt {
					newItems = append(newItems, items[index])
				}
			}
			items = newItems

		}
		json.NewEncoder(w).Encode(map[string]interface{}{"method": "DELETE method", "message": "successfully deleted item"})
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
