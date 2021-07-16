package main

import (
	"fmt"
	"os"
)

type file struct {
	path string
	fs   *os.File
}

func (fs *file) createFile(p string) {
	fs.path = p
	fmt.Printf("Creating file - %s\n", fs.path)
	f, err := os.Create(fs.path)
	if err != nil {
		panic(err)
	}
	fs.fs = f
}

func (fs *file) writeFile() {
	fmt.Println("writing")
	fmt.Fprintln(fs.fs, "data")
}

func (fs *file) closeFile() {
	fmt.Println("closing")
	err := fs.fs.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	fs := &file{}
	fs.createFile("defer.txt")
	defer fs.closeFile()
	fs.writeFile()
}
