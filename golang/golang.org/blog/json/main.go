package main

import (
	"encoding/json"
	"fmt"
)

// Link: https://blog.golang.org/json

type Message struct {
	Name string
	Body string
	Time int64
}

type FamilyMember struct {
	Name    string
	Age     float64
	Parents []string
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	// How does Unmarshal identify the fields in which to store the decoded data?
	// For a given JSON key "Foo", Unmarshal will look through
	// the destination struct's fields to find (in order of preference):
	// An exported field with a tag of "Foo" (see the Go spec for more on struct tags),
	// An exported field named "Foo", or
	// An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
	m = *new(Message)

	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", m)
	m = *new(Message)

	b = []byte(`{"Name":"Bob","Food":"Pickle"}`)
	// Unmarshal will decode only the fields that it can find in the destination type.
	// In this case, only the Name field of m will be populated, and the Food field will be ignored.
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", m)
	m = *new(Message)

	b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	fmt.Printf("%v\n", f)

	mm := f.(map[string]interface{})
	for k, v := range mm {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	var mp FamilyMember
	err = json.Unmarshal(b, &mp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", mp)
}
