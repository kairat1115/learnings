package main

import (
	"bufio"
	"compress/lzw"
	"fmt"
	"io"
	"os"
)

// Link: https://blog.golang.org/gif-decoder

type reader interface {
	io.Reader
	io.ByteReader
}

type blockReader struct {
	r     reader
	slice []byte
	tmp   [256]byte
}

func (b *blockReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if len(b.slice) == 0 {
		blockLen, err := b.r.ReadByte()
		if err != nil {
			return 0, err
		}
		if blockLen == 0 {
			return 0, io.EOF
		}
		b.slice = b.tmp[0:blockLen]
		if _, err := io.ReadFull(b.r, b.slice); err != nil {
			return 0, err
		}
	}
	n := copy(p, b.slice)
	b.slice = b.slice[n:]
	return n, nil
}

var litWidth = 8

func main() {
	f, err := os.Open("sergi-animated.gif")
	if err != nil {
		panic(err)
	}
	imageFile := bufio.NewReader(f)

	deblockingReader := &blockReader{r: imageFile}

	buf := make([]byte, 256)

	lzwr := lzw.NewReader(deblockingReader, lzw.LSB, int(litWidth))
	if _, err := io.ReadFull(lzwr, buf); err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
}
