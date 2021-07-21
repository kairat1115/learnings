package main

import "fmt"

type A struct {
	val int
}

type B struct {
	val int
}

func (a A) a1() {
	a.val++
}

func (a *A) a2() {
	a.val++
}

func (b B) b1() {
	b.val++
}

func (b *B) b2() {
	(*b).val++ // b.val converts to (*b).val
}

func main() {
	a := A{1}
	b := &B{1}
	fmt.Println(a.val, b.val)
	a.a1()
	b.b1()
	fmt.Println(a.val, b.val)
	(&a).a2() // a.a2() converts to (&a).a2()
	b.b2()
	fmt.Println(a.val, b.val)
}
