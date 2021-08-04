package main

// Link: https://blog.golang.org/http-tracing

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	trace := &httptrace.ClientTrace{
		DNSDone: func(di httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", di)
		},
		GotConn: func(gci httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", gci)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
}
