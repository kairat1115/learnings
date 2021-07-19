package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct{}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string {
	return "Designing patterns!"
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	v, err := time.Parse(time.RubyDate, s)
	if err != nil {
		return err
	}

	*t = Timestamp(v)
	return nil
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}, Llama{}, JavaProgrammer{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	input := `
	{
		"created_at": "Thu May 31 00:00:01 +0001 2012"
	}`

	// var val map[string]interface{}
	var val map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v), time.Time(v))
	}
}
