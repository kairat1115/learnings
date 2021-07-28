package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Link: https://blog.golang.org/gob

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	if err := enc.Encode(P{3, 4, 5, "Pythagoras"}); err != nil {
		log.Fatal("encode error:", err)
	}

	var q Q
	if err := dec.Decode(&q); err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
