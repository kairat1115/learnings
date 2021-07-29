package main

import (
	"fmt"
	"reflect"
)

// Link: https://blog.golang.org/laws-of-reflection

type myInt int

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var z myInt = 7
	// Kind retrieves underlying type of variable
	// No matter if variable assigned to myInt type, the underlying is int
	fmt.Println("kind of myInt:", reflect.ValueOf(z).Kind()) // int

	// get interface of reflect.value and type-assert to float64
	y := v.Interface().(float64)
	fmt.Println("interface value:", y)
	fmt.Println("interface value (no assert):", v.Interface())

	// we can use returned value in format string
	fmt.Printf("value is %7.1e\n", v.Interface())

	// check if v can be setted
	fmt.Println("settability of v:", v.CanSet())
	// v.SetFloat(7.1) // Will panic, because it passed-by-value at line 15.

	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	// will write false, because it prints if p can be settable
	fmt.Println("settability of p:", p.CanSet())

	v = p.Elem()
	// now it can set, because we indirects through the pointer to real value
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface()) // 7.1
	fmt.Println(x)             // 7.1

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		// returns i'th field of struct T
		f := s.Field(i)
		// typeOfT.Field(i).Name - get a name of i'th field of struct T
		// f.Type() - get type of i'th field of struct T
		// f.Interface() - get interface of i'th field of struct T
		// f.Interface() will return value
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	// since s is like a pointer, we can set some values
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
