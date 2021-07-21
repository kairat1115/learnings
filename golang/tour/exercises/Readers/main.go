package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = byte('A') // 0x41
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
