package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router

	//Items    []Item
	Database *Database
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
		//Items:    []Item{},
		Database: NewDatabase(),
	}
	server.routes()

	return server
}

func (server *Server) routes() {
	v1 := server.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/items", server.listItems()).Methods("GET")
	v1.HandleFunc("/items", server.addItem()).Methods("POST")
	v1.HandleFunc("/items/{id}", server.removeItem()).Methods("DELETE")
}

func (server *Server) addItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil { // Read the request body
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		item.ID = uuid.New()
		//server.Items = append(server.Items, item)

		w.Header().Set("Content-Type", "application/json")
		if err := server.Database.AddItem(item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Added item", item)

		if err := json.NewEncoder(w).Encode(item); err != nil { // Write the response body
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (server *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//items := server.Items
		items, err := server.Database.GetItems()

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(items); err != nil { // Write the response body
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (server *Server) removeItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := mux.Vars(r)["id"]
		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//for i, item := range server.Items {
		items, err := server.Database.GetItems()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, item := range items {
			if item.ID == id {
				if err := server.Database.RemoveItem(idString); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					fmt.Println("Removed item", item)
				}
				return
			}
		}
	}
}
