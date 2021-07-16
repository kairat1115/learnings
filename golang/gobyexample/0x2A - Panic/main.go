package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
}
