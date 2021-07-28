package main

import (
	"fmt"
	"io"
	"os"
)

// Link: https://blog.golang.org/defer-panic-and-recover

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	// We can rewrite this with simple "return io.Copy(dst, src)" since it returns int64 and error
	// written, err = io.Copy(dst, src)

	// BUG: If creation fails, src file remaines opened
	// We should use defer.
	// dst.Close()
	// src.Close()

	// return
	return io.Copy(dst, src)
}

func a() {
	i := 0
	// defer evaluate with a given value at time defer was called
	// so output will be 0, not 1.
	defer fmt.Println(i)
	i++
	return
}

func b() {
	// just new line
	defer fmt.Println()
	// Defer works like a stack
	// first in, last out
	for i := 0; i < 4; i++ {
		// defer will put values like that
		// 0 -> [] -> print
		// 1 -> [0] -> print
		// 2 -> [1, 0] -> print
		// 3 -> [2, 1, 0] -> print
		// nil -> [3, 2, 1, 0] -> print
		// printing 3,2,1,0
		// Outputs 3210
		defer fmt.Print(i)
	}
}

func c() (i int) {
	// a deferred function increments the return value i
	// after the surrounding function returns.
	// Thus, this function returns 2
	defer func() { i++ }()
	return 1
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	// Recursive call does not defers until returned.
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func main() {
	a()
	b()
	fmt.Println(c())

	f()
	fmt.Println("Returned normally from f.")
}
