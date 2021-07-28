package main

import (
	"fmt"
	"strings"
)

// Header: Go Slices: usage and internals
// Link: https://blog.golang.org/slices-intro

func main() {
	fmt.Println(strings.Repeat("=", 10))

	// if you specify count between [] (... is auto count elements)
	// currently [...] is [2]
	// then it means it is array
	a1 := [...]string{"Penn", "Teller"}
	fmt.Println(a1)

	fmt.Println(strings.Repeat("=", 10))

	// this is slice, because count is empty
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	fmt.Println(strings.Repeat("=", 10))

	ingtegers := make(
		[]int, // data type
		3,     // len() - length
		3,     // cap() - capacity
	)
	fmt.Println(ingtegers)

	fmt.Println(strings.Repeat("=", 10))

	// capacity is length from original array[n] to end of array
	// arr := []int{1, 2, 3, 4, 5, 6}
	// ex: b := arr[1:4]
	// b is [2, 3, 4], len(b) is 3, cap(b) is 5
	// c := b[1:2]
	// c is [3], len(b) is 1, cap(b) is 4 because it will count from b[0:len(arr)] original [3, 4, 5, 6]
	arr := []int{1, 2, 3, 4, 5, 6}
	a2 := arr[1:4]
	fmt.Println(a2, len(a2), cap(a2))
	b2 := a2[1:2]
	fmt.Println(b2, len(b2), cap(b2))

	fmt.Println(strings.Repeat("=", 10))

	a3 := []string{"John", "Paul"}
	b3 := []string{"George", "Ringo", "Pete"}
	a3 = append(a3, b3...) // equivalent to "append(a3, b3[0], b3[1], b3[2])"
	// a3 == []string{"John", "Paul", "George", "Ringo", "Pete"}
	fmt.Println(a3)

	fmt.Println(strings.Repeat("=", 10))
}
