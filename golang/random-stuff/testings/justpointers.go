package main

import "fmt"

func main() {
	asdf := 123
	fmt.Println(asdf)
	p1 := &asdf
	fmt.Println(p1, *p1)
	p2 := &p1
	fmt.Println(p2, *p2, **p2)
	p3 := &p2
	fmt.Println(p3, *p3, **p3, ***p3)
	p4 := &p3
	fmt.Println(p4, *p4, **p4, ***p4, ****p4)
}
