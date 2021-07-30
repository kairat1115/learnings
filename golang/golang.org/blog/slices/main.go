package main

// Link: https://blog.golang.org/slices

import (
	"bytes"
	"fmt"
	"unicode"
)

type sliceHeader struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

type path []byte

func (this *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*this, []byte("/"))
	if i > -1 {
		*this = (*this)[0:i]
	}
}

func (this path) ValTruncateAtFinalSlash() {
	i := bytes.LastIndex(this, []byte("/"))
	if i > -1 {
		this = this[0:i]
	}
}

func (this path) ToUpper() {
	for i, b := range this {
		if 'a' <= b && b <= 'z' {
			this[i] = b + 'A' - 'a'
		}
	}
}

func (this *path) PtrToUpper() {
	for i, b := range *this {
		if 'a' <= b && b <= 'z' {
			(*this)[i] = b + 'A' - 'a'
		}
	}
}

type unipath []rune

func (this unipath) ToUpper() {
	for i, b := range this {
		this[i] = unicode.ToUpper(b)
	}
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {
	var buffer [256]byte
	{
		var slice []byte = buffer[100:150]
		fmt.Println(slice, len(slice), cap(slice))
	}
	{
		slice := buffer[100:150]
		fmt.Println(slice, len(slice), cap(slice))
		slice2 := slice[5:10]
		fmt.Println(slice2, len(slice2), cap(slice2))
		slice = slice[5:10]
		fmt.Println(slice, len(slice), cap(slice))
		slice = slice[1 : len(slice)-1]
		fmt.Println(slice, len(slice), cap(slice))
	}
	{
		// buffer[100:150]
		slice := sliceHeader{
			Length:        50,
			ZerothElement: &buffer[100],
		}
		fmt.Println(slice)
		// slice[5:10]
		slice = sliceHeader{
			Length:        5,
			ZerothElement: &buffer[105],
		}
		fmt.Println(slice)
		// slice[1 : len(slice)-1]
		slice = sliceHeader{
			Length:        3,
			ZerothElement: &buffer[106],
		}
		fmt.Println(slice)
	}
	{
		slice := buffer[100:150]
		slashPos := bytes.IndexRune(slice, '/')
		fmt.Println(slashPos)
	}
	{
		slice := buffer[10:20]
		for i := 0; i < len(slice); i++ {
			slice[i] = byte(i)
		}
		fmt.Println("before", slice)
		AddOneToEachElement(slice)
		fmt.Println("after", slice)
	}
	{
		slice := buffer[50:100]
		fmt.Println("Before: len(slice)    =", len(slice))
		newSlice := SubtractOneFromLength(slice)
		fmt.Println("After:  len(slice)    =", len(slice))
		fmt.Println("After:  len(newSlice) =", len(newSlice))

		fmt.Println("Before: len(slice) =", len(slice))
		PtrSubtractOneFromLength(&slice)
		fmt.Println("After:  len(slice) =", len(slice))
	}
	{
		pathName := path("/usr/bin/tso")
		pathName.TruncateAtFinalSlash()
		fmt.Printf("%s\n", pathName)
	}
	{
		// Exercise.
		pathName := path("/usr/bin/tso")
		// pathName []byte passed by value to function, so it wont be modified.
		// in function we replacing original with new slice
		// but we can change it, because it still has pointer to underlying array.
		pathName.ValTruncateAtFinalSlash()
		fmt.Printf("%s\n", pathName)
	}
	{
		pathName := path("/usr/bin/tso")
		pathName.ToUpper()
		fmt.Printf("%s\n", pathName)
		// Exercise with ptr
		// No changes
		pathName = path("/usr/bin/tso")
		pathName.PtrToUpper()
		fmt.Printf("%s\n", pathName)
	}
	{
		// Advanced exercise
		pathName := unipath("/путь/до/файла")
		pathName.ToUpper()
		fmt.Printf("%s\n", string(pathName))
	}
	{
		var iBuffer [10]int
		slice := iBuffer[0:0]
		for i := 0; i < 20; i++ {
			slice = Extend(slice, i)
			fmt.Println(slice)
		}
	}
}
