package main

import (
	_ "github.com/go-sql-driver/mysql"
	"http-server/api"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", api.NewServer()); err != nil {
		panic(err)
	}
}
