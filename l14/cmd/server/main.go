package main

import (
	"hw15/internal/server"
	"net/http"
)

func main() {
	server := server.Server{}
	mux := server.GetRouter()
	http.ListenAndServe(":8000", mux)
}
