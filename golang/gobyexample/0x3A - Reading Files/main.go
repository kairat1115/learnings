package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var seek_value int64
	switch runtime.GOOS {
	case "windows":
		seek_value = 7
	default:
		seek_value = 6
	}
	// 6 - for linux, 7 - for windows
	// linux has \n, but windows has \r\n

	dat, err := ioutil.ReadFile("dat.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open("dat.txt")
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(seek_value, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(seek_value, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", b4)
}
