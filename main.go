package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"http-server/api"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}).Handler(api.NewServer())))
}
