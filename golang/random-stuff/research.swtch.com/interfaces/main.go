package main

import (
	"strconv"
)

// Link: https://research.swtch.com/interfaces

type ReadCloser interface {
	Read(b []byte) (n int, err error)
	Close()
}

func ReadAndClose(r ReadCloser, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}

type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'E', -1, 64)
	}
	return "???"
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	// no code :(
}
