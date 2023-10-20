package main

import (
	"http-server/api"
	"net/http"
)

func main() {
	server := api.NewServer()
	http.ListenAndServe(":80", server)
}