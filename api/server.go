package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

var path = "/api/v1"

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router

	Items []Item
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
		Items:  []Item{},
	}
	server.routes()
	return server
}

func (server *Server) routes() {
	server.HandleFunc(path+"/items", server.listItems()).Methods("GET")
	server.HandleFunc(path+"/items", server.addItem()).Methods("POST")
	server.HandleFunc(path+"/items{id}", server.removeItem()).Methods("DELETE")
}

func (server *Server) addItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item Item
		if err := json.NewDecoder(r.Body).Decode(&item); err != nil { // Read the request body
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		item.ID = uuid.New()
		server.Items = append(server.Items, item)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil { // Write the response body
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (server *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(server.Items); err != nil { // Write the response body
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

		for i, item := range server.Items {
			if item.ID == id {
				server.Items = append(server.Items[:i], server.Items[i+1:]...)
				break
			}
		}
	}
}
