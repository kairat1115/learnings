package main

import (
	server "learn-go-with-tests/http-server"
	"log"
	"net/http"
)

func main() {
	store := server.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
