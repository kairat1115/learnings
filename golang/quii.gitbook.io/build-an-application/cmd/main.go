package main

import (
	server "learn-go-with-tests/build-an-application"
	"log"
	"net/http"
)

func main() {
	store := server.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
