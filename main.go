package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"http-server/api"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080",
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"POST", "GET", "DELETE", "PUT", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"*"}))(api.NewServer()))))
}
