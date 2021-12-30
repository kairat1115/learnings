package main

import (
	"net/http"
	q_and_a "q-and-a-httphandlersrevisited"
)

func main() {
	mongoService := q_and_a.NewMongoUserService()
	server := q_and_a.NewUserServer(mongoService)
	http.ListenAndServe(":8000", http.HandlerFunc(server.RegisterUser))
}
