package main

// Link: https://blog.golang.org/constants

import "fmt"

func main() {
	// This is untyped string constant
	const hello = "Hello, 世界"

	// This is typed string constant
	// because given type "string"
	const typedHello string = "Hello, 世界"

	// if a variable is typed, it will obey to rules of go.
	// meaning, it cannot be assigned to other types other than string

	// this will work
	var s string
	s = typedHello
	fmt.Println(s)

	type MyString string
	var m MyString

	// but this will not.
	// cannot use typedHello (constant "Hello, 世界" of type string) as MyString value in assignment
	// m = typedHello

	// only same type is allowed
	const myStringHello MyString = "Hello, 世界"
	m = myStringHello
	// but, we can convert
	m = MyString(typedHello)
	fmt.Println(m)

	// cool trick
	// get maximum number for type
	// e.g. get max value for uint
	// we just invert zeroed value of uint
	// 0 becomes 1
	const MaxUint = ^uint(0)
	fmt.Printf("%x\n", MaxUint)
}
