package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	var rot byte = 13

	n, err = r.r.Read(p)
	for i := range p {
		if p[i] >= byte('A') && p[i] <= byte('Z') {
			p[i] = (p[i]-byte('A')+rot)%26 + byte('A')
		} else if p[i] >= byte('a') && p[i] <= byte('z') {
			p[i] = (p[i]-byte('a')+rot)%26 + byte('a')
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
